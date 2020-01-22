package eth

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"log"
	"math/big"
	"os"
	"regexp"
	"strings"
	"time"

	"gopkg.in/eapache/go-resiliency.v1/retrier"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/javiercbk/ppv-crypto/server/cryptocurrency/eth/ppvevent"
)

// InvalidTimeErr is returned when trying to deploy a contract with wrong start or end time
type InvalidTimeErr string

func (e InvalidTimeErr) Error() string {
	return string(e)
}

// InvalidPriceErr is thrown when trying to deploy a contract with wrong price
type InvalidPriceErr string

func (e InvalidPriceErr) Error() string {
	return string(e)
}

// PPVContractUpdater handles updates in a PPVContract
type PPVContractUpdater interface {
	ChangePrice(ctx context.Context, eventID int64, newPrice int64, block BlockMetadata) error
	NewSubscription(ctx context.Context, eventID int64, newSubscription NewSubscription, block BlockMetadata) error
	NewUnsubscription(ctx context.Context, eventID int64, newUnsubscription NewUnsubscription, block BlockMetadata) error
}

const (

	// LogPPVEventStartedHex is the keccak256("PPVEventStarted()")"
	LogPPVEventStartedHex = "0e10146934a460c3f27a8ca4c63fdad156eed0ae2051edc0c1839268ade6c9ad"
	// LogPPVEventEndedHex us the keccak256("PPVEventEnded()")
	LogPPVEventEndedHex = "545c09a6805d0d5edc24946e122843a3b3cfb83cef050826f143983353509e2f"
	// LogPriceChangedHex is the keccak256("PriceChanged(uint256)")
	LogPriceChangedHex = "a6dc15bdb68da224c66db4b3838d9a2b205138e8cff6774e57d0af91e196d622"
	// LogNewSubscriptionHex is the keccak256("NewSubscription(address,uint256)")
	LogNewSubscriptionHex = "1e05df24f73db39faa0c2d5d26727d08632debce09833123a69214ba943e07c2"
	// LogNewUnsubscriptionHex is the keccak256("NewUnsubscription(address,uint256)")
	LogNewUnsubscriptionHex = "79d9d14aed97c97b8dc663528ec3b03eecb8d467956fc03f37179b188a04efa4"
	// ErrInvalidTime is returned when trying to deploy a contract with wrong start or end time
	ErrInvalidTime InvalidTimeErr = "contract time is invalid"
	// ErrInvalidPrice is thrown when trying to deploy a contract with wrong price
	ErrInvalidPrice InvalidPriceErr = "contract price is invalid"
)

var (
	logPPVEventStartedHex   = common.HexToHash(LogPPVEventStartedHex)
	logPPVEventEndedHex     = common.HexToHash(LogPPVEventEndedHex)
	logPriceChangedHex      = common.HexToHash(LogPriceChangedHex)
	logNewSubscriptionHex   = common.HexToHash(LogNewSubscriptionHex)
	logNewUnsubscriptionHex = common.HexToHash(LogNewUnsubscriptionHex)
)

// PriceChanged emmited on price changed
type PriceChanged struct {
	PPVEventPrice *big.Int
}

// NewSubscription emmited on new subscription
type NewSubscription struct {
	Subscriptor common.Address
	Price       *big.Int
}

// NewUnsubscription emmited on new unsubscription
type NewUnsubscription struct {
	Subscriptor common.Address
	Price       *big.Int
}

// Validates an ethereum address
var addressRE = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

// Connect connects to the ethereum network
func Connect(networkURL string) (*ethclient.Client, error) {
	return ethclient.Dial(networkURL)
}

// ValidateAddress validates an ethereum address
func ValidateAddress(address string) bool {
	return addressRE.MatchString(address)
}

// NeedsReconnect determinates if an error is due a connection quitting
func NeedsReconnect(err error) bool {
	return errors.Is(err, rpc.ErrClientQuit)
}

// BlockMetadata contains information about the block where a transaction happened
type BlockMetadata struct {
	BlockHash   common.Hash
	BlockNumber uint64
	TxHash      common.Hash
	TxNumber    uint64
}

// ContractWatcher is able to watch a events in smart contracts
type ContractWatcher struct {
	contractAddress common.Address
	eventID         int64
	client          *ethclient.Client
	logger          *log.Logger
	contract        *ppvevent.Ppvevent
	updater         PPVContractUpdater
}

// SubscribeToEvents given a contract address it will read every event until the context is cancelled
func (w ContractWatcher) SubscribeToEvents(ctx context.Context) error {
	r := retrier.New(retrier.ConstantBackoff(3, 100*time.Millisecond), nil)
	contractAddress := w.contractAddress
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(ppvevent.PpveventABI)))
	if err != nil {
		w.logger.Printf("error reading smart contract ABI: %v\n", err)
		return err
	}
	logs := make(chan types.Log)
	defer close(logs)
	sub, err := w.client.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		w.logger.Printf("error subscribing to contracts logs %v\n", err)
	}
	defer sub.Unsubscribe()
	done := ctx.Done()
	for {
		select {
		case err := <-sub.Err():
			w.logger.Printf("error in subscription %v\n", err)
			return err
		case contractLog := <-logs:
			switch contractLog.Topics[0].Hex() {
			case LogPPVEventEndedHex:
				// on event end just stop watching this contract
				return nil
			case LogPriceChangedHex:
				// on price change update the DB
				var priceChangeEvent PriceChanged
				err := contractAbi.Unpack(&priceChangeEvent, "PriceChanged", contractLog.Data)
				if err != nil {
					w.logger.Printf("error unpacking PriceChange event from subscription: %v\n", err)
					// TODO: store an error signal or some kind of data to re-process this event update
				}
				err = r.RunCtx(ctx, func(retrierCtx context.Context) error {
					return w.updater.ChangePrice(retrierCtx, w.eventID, priceChangeEvent.PPVEventPrice.Int64(), toBlockMetadata(contractLog))
				})
				if err != nil {
					w.logger.Printf("error changing price %v\n", err)
					// TODO: handle the case where the change price failed three time
				}
			case LogNewSubscriptionHex:
				var newSubscriptionEvent NewSubscription
				err := contractAbi.Unpack(&newSubscriptionEvent, "NewSubscription", contractLog.Data)
				if err != nil {
					w.logger.Printf("error unpacking NewSubscription event from subscription: %v\n", err)
					// TODO: store an error signal or some kind of data to re-process this event update
				}
				err = r.RunCtx(ctx, func(retrierCtx context.Context) error {
					return w.updater.NewSubscription(retrierCtx, w.eventID, newSubscriptionEvent, toBlockMetadata(contractLog))
				})
				if err != nil {
					w.logger.Printf("error registering subscription %v\n", err)
					// TODO: queue the error somewhere to reprocess this event
				}
			case LogNewUnsubscriptionHex:
				var newUnsubscriptionEvent NewUnsubscription
				err := contractAbi.Unpack(&newUnsubscriptionEvent, "NewUnsubscription", contractLog.Data)
				if err != nil {
					w.logger.Printf("error unpacking NewUnsubscription event from subscription: %v\n", err)
					// TODO: store an error signal or some kind of data to re-process this event update
				}
				err = r.RunCtx(ctx, func(retrierCtx context.Context) error {
					return w.updater.NewUnsubscription(retrierCtx, w.eventID, newUnsubscriptionEvent, toBlockMetadata(contractLog))
				})
				if err != nil {
					w.logger.Printf("error registering unsubscription %v\n", err)
					// TODO: queue the error somewhere to reprocess this event
				}
			default:
				// on PPVEventStarted do nothing
			}
		case <-done:
			// if context is done then end the function execution
			return nil
		}
	}
}

func toBlockMetadata(contractLog types.Log) BlockMetadata {
	return BlockMetadata{
		BlockHash:   contractLog.BlockHash,
		BlockNumber: contractLog.BlockNumber,
		TxHash:      contractLog.TxHash,
		TxNumber:    uint64(contractLog.TxIndex),
	}
}

// NewContractWatcher creates a new contract watcher
func NewContractWatcher(client *ethclient.Client, logger *log.Logger, contractAddress string, eventID int64, updater PPVContractUpdater) (ContractWatcher, error) {
	address := common.HexToAddress(contractAddress)
	contract, err := ppvevent.NewPpvevent(address, client)
	return ContractWatcher{
		contractAddress: address,
		eventID:         eventID,
		client:          client,
		contract:        contract,
		logger:          logger,
		updater:         updater,
	}, err
}

// DeployedContract contains all the information of a recently deployed contract
type DeployedContract struct {
	Address  common.Address
	Tx       types.Transaction
	PPVEvent *ppvevent.Ppvevent
}

// ProspectPPVEvent contains all the information to deploy a ppv smart contract
type ProspectPPVEvent struct {
	Start time.Time
	End   time.Time
	Price int64
}

// SmartContractDeployer can deploy smart contracts
type SmartContractDeployer struct {
	privateKey  *ecdsa.PrivateKey
	logger      *log.Logger
	client      *ethclient.Client
	fromAddress common.Address
}

// DeployNewPPVSmartContract deploys a new PPV smart contract
func (d *SmartContractDeployer) DeployNewPPVSmartContract(ctx context.Context, prospectEvent ProspectPPVEvent, deployedContract *DeployedContract) error {
	if prospectEvent.Start.IsZero() || prospectEvent.End.IsZero() || prospectEvent.Start.After(prospectEvent.End) {
		return ErrInvalidTime
	}
	if prospectEvent.Price <= 0 {
		return ErrInvalidPrice
	}
	nonce, err := d.client.PendingNonceAt(ctx, d.fromAddress)
	if err != nil {
		d.logger.Printf("error getting PendingNonceAt %v", err)
		return err
	}
	gasPrice, err := d.client.SuggestGasPrice(ctx)
	if err != nil {
		d.logger.Printf("error getting SuggestGasPrice %v", err)
		return err
	}
	auth := bind.NewKeyedTransactor(d.privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = 0          // in units
	auth.GasPrice = gasPrice
	auth.Context = ctx
	eventStart := big.NewInt(prospectEvent.Start.Unix())
	eventEnd := big.NewInt(prospectEvent.End.Unix())
	eventPrice := big.NewInt(prospectEvent.Price)
	newContractAddress, transaction, instance, err := ppvevent.DeployPpvevent(auth, d.client, eventStart, eventEnd, eventPrice)
	deployedContract.Address = newContractAddress
	deployedContract.Tx = *transaction
	deployedContract.PPVEvent = instance
	if err != nil {
		d.logger.Printf("failed to deploy smart contract %v", err)
	}
	return err
}

// NewSmartContractDeployer creates a SmartContractDeployer
func NewSmartContractDeployer(client *ethclient.Client, logger *log.Logger, privateKeyBytes []byte) (*SmartContractDeployer, error) {
	deployer := &SmartContractDeployer{}
	// removes the 0x from the private key bytes
	privateKey, err := crypto.HexToECDSA(string(privateKeyBytes[2:]))
	if err != nil {
		logger.Printf("error transforming private key to ECDSA: %v", err)
		os.Exit(1)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
		return deployer, err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	deployer.privateKey = privateKey
	deployer.logger = logger
	deployer.client = client
	deployer.fromAddress = fromAddress
	return deployer, nil
}
