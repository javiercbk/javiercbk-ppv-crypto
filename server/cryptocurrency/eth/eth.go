package cryptocurrency

import (
	"context"
	"errors"
	"log"
	"math/big"
	"regexp"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/javiercbk/ppv-crypto/server/cryptocurrency"
	"github.com/javiercbk/ppv-crypto/server/cryptocurrency/eth/smartcontract"
	"github.com/javiercbk/ppv-crypto/server/event"
	"github.com/javiercbk/ppv-crypto/server/models"
	"github.com/volatiletech/null"
)

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

// ContractWatcher is able to watch a events in smart contracts
type ContractWatcher struct {
	contractAddress string
	eventID         int64
	client          *ethclient.Client
	logger          *log.Logger
	contract        *smartcontract.Smartcontract
	eventAPI        event.API
}

// SubscribeToEvents given a contract address it will read every event until the context is cancelled
func (w ContractWatcher) SubscribeToEvents(ctx context.Context) error {
	contractAddress := common.HexToAddress(w.contractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(smartcontract.SmartcontractABI)))
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
				ppvEvent := &models.PayPerViewEvent{
					ID:       w.eventID,
					PriceETH: null.Int64From(priceChangeEvent.PPVEventPrice.Int64()),
				}
				err = w.eventAPI.UpdateEvent(ctx, ppvEvent)
				if err != nil {
					w.logger.Printf("error updating event price: %v\n", err)
					// TODO: store an error signal or some kind of data to re-process this event update
				}
			case LogNewSubscriptionHex:
				var newSubscriptionEvent NewSubscription
				err := contractAbi.Unpack(&newSubscriptionEvent, "NewSubscription", contractLog.Data)
				if err != nil {
					w.logger.Printf("error unpacking NewSubscription event from subscription: %v\n", err)
					// TODO: store an error signal or some kind of data to re-process this event update
				}
				payment := event.PPVEventPayment{
					EventID:        w.eventID,
					WalletAddress:  newSubscriptionEvent.Subscriptor.Hex(),
					Price:          newSubscriptionEvent.Price.Int64(),
					CryptoCurrency: cryptocurrency.ETH,
					BlockHash:      contractLog.BlockHash.Hex(),
					BlockNumber:    strconv.FormatUint(uint64(contractLog.BlockNumber), 16),
					TxHash:         contractLog.TxHash.Hex(),
					TxNumber:       strconv.FormatUint(uint64(contractLog.TxIndex), 16),
				}
				w.eventAPI.RegisterSubscription(ctx, payment)
			case LogNewUnsubscriptionHex:
				var newUnsubscriptionEvent NewUnsubscription
				err := contractAbi.Unpack(&newUnsubscriptionEvent, "NewUnsubscription", contractLog.Data)
				if err != nil {
					w.logger.Printf("error unpacking NewUnsubscription event from subscription: %v\n", err)
				}
				payment := event.PPVEventPayment{
					EventID:        w.eventID,
					WalletAddress:  newUnsubscriptionEvent.Subscriptor.Hex(),
					Price:          newUnsubscriptionEvent.Price.Int64(),
					CryptoCurrency: cryptocurrency.ETH,
					BlockHash:      contractLog.BlockHash.Hex(),
					BlockNumber:    strconv.FormatUint(uint64(contractLog.BlockNumber), 16),
					TxHash:         contractLog.TxHash.Hex(),
					TxNumber:       strconv.FormatUint(uint64(contractLog.TxIndex), 16),
				}
				w.eventAPI.RegisterUnsubscription(ctx, payment)
			default:
				// on PPVEventStarted do nothing
			}
		case <-done:
			// if context is done then end the function execution
			return nil
		}
	}
}

// NewContractWatcher creates a new contract watcher
func NewContractWatcher(client *ethclient.Client, eventAPI event.API, logger *log.Logger, contractAddress string, eventID int64) ContractWatcher {
	return ContractWatcher{
		contractAddress: contractAddress,
		eventID:         eventID,
		client:          client,
		contract:        smartcontract.NewSmartcontract(common.HexToAddress(contractAddress)),
		logger:          logger,
		eventAPI:        eventAPI,
	}
}
