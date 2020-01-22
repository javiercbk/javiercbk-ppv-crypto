package event

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/javiercbk/ppv-crypto/server/cryptocurrency"
	"github.com/javiercbk/ppv-crypto/server/cryptocurrency/eth"
	"github.com/javiercbk/ppv-crypto/server/http/security"
	"github.com/javiercbk/ppv-crypto/server/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// InexistentEventErr is returned when attempting to retrieve a non existing event
type InexistentEventErr string

func (e InexistentEventErr) Error() string {
	return string(e)
}

// InexistentPaymentErr is returned when attempting to retrieve a non existing payment
type InexistentPaymentErr string

func (e InexistentPaymentErr) Error() string {
	return string(e)
}

const (
	// ErrInexistentEvent is returned when attempting to retrieve a non existing event
	ErrInexistentEvent InexistentEventErr = "inexistent event"
	// ErrInexistentPayment is returned when attempting to retrieve a non existing payment
	ErrInexistentPayment InexistentPaymentErr = "inexistent payment"
)

// APIFactory is a function capable of creating an Event API
type APIFactory func(logger *log.Logger, db *sql.DB, deployer *eth.SmartContractDeployer) API

// PPVEventAndSubscription an event containing the user payments
type PPVEventAndSubscription struct {
	models.PayPerViewEvent
	Payments []models.Payment `json:"payments"`
}

// NullablePayment is the same as model.Payment but all their fields are nullable
type nullablePayment struct {
	PID                            null.Int64  `boil:"id" json:"id" toml:"id" yaml:"id"`
	PaymentUserID                  null.Int64  `boil:"user_id" json:"userID,omitempty" toml:"userID" yaml:"userID,omitempty"`
	PaymentPayPerViewEventID       null.Int64  `boil:"pay_per_view_event_id" json:"payPerViewEventID" toml:"payPerViewEventID" yaml:"payPerViewEventID"`
	PaymentCurrency                null.String `boil:"currency" json:"currency" toml:"currency" yaml:"currency"`
	PaymentCurrencyPaymentID       null.String `boil:"currency_payment_id" json:"currencyPaymentID,omitempty" toml:"currencyPaymentID" yaml:"currencyPaymentID,omitempty"`
	PaymentAmount                  null.Int64  `boil:"amount" json:"amount,omitempty" toml:"amount" yaml:"amount,omitempty"`
	PaymentWalletAddress           null.String `boil:"wallet_address" json:"walletAddress,omitempty" toml:"walletAddress" yaml:"walletAddress,omitempty"`
	PaymentStatus                  null.String `boil:"status" json:"status" toml:"status" yaml:"status"`
	PaymentBlockHash               null.String `boil:"block_hash" json:"blockHash,omitempty" toml:"blockHash" yaml:"blockHash,omitempty"`
	PaymentBlockNumberHex          null.String `boil:"block_number_hex" json:"blockNumberHex,omitempty" toml:"blockNumberHex" yaml:"blockNumberHex,omitempty"`
	PaymentTXHash                  null.String `boil:"tx_hash" json:"txHash,omitempty" toml:"txHash" yaml:"txHash,omitempty"`
	PaymentTXNumberHex             null.String `boil:"tx_number_hex" json:"txNumberHex,omitempty" toml:"txNumberHex" yaml:"txNumberHex,omitempty"`
	PaymentCancelledBlockHash      null.String `boil:"cancelled_block_hash" json:"cancelledBlockHash,omitempty" toml:"cancelledBlockHash" yaml:"cancelledBlockHash,omitempty"`
	PaymentCancelledBlockNumberHex null.String `boil:"cancelled_block_number_hex" json:"cancelledBlockNumberHex,omitempty" toml:"cancelledBlockNumberHex" yaml:"cancelledBlockNumberHex,omitempty"`
	PaymentCancelledTXHash         null.String `boil:"cancelled_tx_hash" json:"cancelledTXHash,omitempty" toml:"cancelledTXHash" yaml:"cancelledTXHash,omitempty"`
	PaymentCancelledTXNumberHex    null.String `boil:"cancelled_tx_number_hex" json:"cancelledTXNumberHex,omitempty" toml:"cancelledTXNumberHex" yaml:"cancelledTXNumberHex,omitempty"`
	PaymentCancelledAt             null.Time   `boil:"cancelled_at" json:"cancelledAt,omitempty" toml:"cancelledAt" yaml:"cancelledAt,omitempty"`
	PaymentCreatedAt               null.Time   `boil:"created_at" json:"createdAt,omitempty" toml:"createdAt" yaml:"createdAt,omitempty"`
	PaymentUpdatedAt               null.Time   `boil:"updated_at" json:"updatedAt,omitempty" toml:"updatedAt" yaml:"updatedAt,omitempty"`
}

// ppvEventAndSubscriptionResult is used to bind the query results from a query
type ppvEventAndSubscriptionResult struct {
	models.PayPerViewEvent `boil:"pay_per_view_events,bind"`
	nullablePayment        `boil:"payments,bind"`
}

// PPVEvent has all the information of a pay per view event
type PPVEvent struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	EventType   string          `json:"eventType"`
	Start       time.Time       `json:"start"`
	End         time.Time       `json:"end"`
	PriceETH    int64           `json:"priceETH"`
	PriceBTC    int64           `json:"priceBTC"`
	PriceXMR    int64           `json:"priceXMR"`
	CreatedAt   time.Time       `json:"createdAt"`
	UpdatedAt   *time.Time      `json:"updatedAt,omitempty"`
	Payment     *models.Payment `json:"payment,omitempty"`
}

// PPVEventQuery contains all the data to query a PPVEvent
type PPVEventQuery struct {
	ID    int64
	Query string    `json:"query,omitempty" validate:"lte=256"`
	From  time.Time `json:"from,omitempty"`
	To    time.Time `json:"to,omitempty" validate:"gtfield=From|eqfield=From"`
}

// PPVEventPayment contains the information to process a payment
type PPVEventPayment struct {
	EventID        int64                         `json:"eventID"`
	Price          int64                         `json:"price"`
	CryptoCurrency cryptocurrency.CryptoCurrency `json:"cryptoCurrency"`
	PaymentID      string                        `json:"paymentID"`
	WalletAddress  string                        `json:"walletAddress"`
	BlockHash      string                        `json:"blockHash"`
	BlockNumber    string                        `json:"blockNumber"`
	TxHash         string                        `json:"txHash"`
	TxNumber       string                        `json:"txNumber"`
	Status         string                        `json:"status"`
	User           security.JWTUser
}

// API is an event API interface
type API interface {
	RetrieveEvents(ctx context.Context, ppvQuery PPVEventQuery, user *security.JWTUser) ([]PPVEventAndSubscription, error)
	CreateEvent(ctx context.Context, ppvEvent *models.PayPerViewEvent) error
	RetrieveEvent(ctx context.Context, eventID int64, user *security.JWTUser) (PPVEventAndSubscription, error)
	UpdateEvent(ctx context.Context, ppvEvent *models.PayPerViewEvent) error
	ProcessPayment(ctx context.Context, paymentData PPVEventPayment, payment *models.Payment) error
	RetrievePaymentStatus(ctx context.Context, eventID int64, user security.JWTUser) (models.PaymentSlice, error)
	RegisterSubscription(ctx context.Context, payment PPVEventPayment) (*models.Payment, error)
	RegisterUnsubscription(ctx context.Context, payment PPVEventPayment) (*models.Payment, error)
}

type api struct {
	logger   *log.Logger
	db       *sql.DB
	deployer *eth.SmartContractDeployer
}

// NewAPI creates a new authentication API
func NewAPI(logger *log.Logger, db *sql.DB, deployer *eth.SmartContractDeployer) API {
	return api{
		logger:   logger,
		db:       db,
		deployer: deployer,
	}
}

// RetrieveEvents retrieves the list of events with or without paymeny information from a user
func (api api) RetrieveEvents(ctx context.Context, ppvQuery PPVEventQuery, user *security.JWTUser) ([]PPVEventAndSubscription, error) {
	events, err := api.retrieveEventsFromDatabase(ctx, ppvQuery, user)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return events, err
	}
	return events, nil
}

// CreateEvent creates a pay per view event
func (api api) CreateEvent(ctx context.Context, ppvEvent *models.PayPerViewEvent) error {
	tx, err := api.db.BeginTx(ctx, nil)
	if err != nil {
		api.logger.Printf("error beginning transaction %v", err)
		return err
	}
	err = ppvEvent.Insert(ctx, tx, boil.Whitelist(
		models.PayPerViewEventColumns.Name,
		models.PayPerViewEventColumns.Description,
		models.PayPerViewEventColumns.EventType,
		models.PayPerViewEventColumns.Start,
		models.PayPerViewEventColumns.End,
		models.PayPerViewEventColumns.PriceEth,
		models.PayPerViewEventColumns.PriceBTC,
		models.PayPerViewEventColumns.PriceXMR,
		models.PayPerViewEventColumns.CreatedAt,
	))
	if err != nil {
		api.logger.Printf("error inserting event: %v", err)
		return err
	}
	if ppvEvent.Start.Valid && ppvEvent.End.Valid && ppvEvent.PriceEth.Valid {
		// Maybe this should be done in a go routine and
		deployedContract := eth.DeployedContract{}
		prospectEvent := eth.ProspectPPVEvent{
			Start: ppvEvent.Start.Time,
			End:   ppvEvent.End.Time,
			Price: ppvEvent.PriceEth.Int64,
		}
		err = api.deployer.DeployNewPPVSmartContract(ctx, prospectEvent, &deployedContract)
		if err != nil {
			defer tx.Rollback()
			api.logger.Printf("error deploying smart contract: %v", err)
			return err
		}
		contractAddress := deployedContract.Address.Hex()
		smartContract := &models.SmartContract{
			PayPerViewEventID: ppvEvent.ID,
			Address:           contractAddress,
			Currency:          cryptocurrency.ETH.String(),
		}
		err = ppvEvent.AddSmartContracts(ctx, tx, true, smartContract)
		if err != nil {
			api.logger.Printf("error inserting smart contract with address '%s', error: %v", contractAddress, err)
			// since the contract is already deployed we do not want to return the error and roll back the whole thing
		}
	}
	err = tx.Commit()
	if err != nil {
		api.logger.Printf("error commiting transaction inserting event: %v", err)
		return err
	}
	return nil
}

func (api api) RetrieveEvent(ctx context.Context, eventID int64, user *security.JWTUser) (PPVEventAndSubscription, error) {
	var eventSubscription PPVEventAndSubscription
	events, err := api.retrieveEventsFromDatabase(ctx, PPVEventQuery{
		ID: eventID,
	}, user)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return eventSubscription, ErrInexistentEvent
		}
		return eventSubscription, err
	}
	eventSubscription = events[0]
	return eventSubscription, err
}

func (api api) UpdateEvent(ctx context.Context, ppvEvent *models.PayPerViewEvent) error {
	_, err := ppvEvent.Update(ctx, api.db, boil.Infer())
	return err
}

// ProcessPayment triggers a reviewing payment for an event
func (api api) ProcessPayment(ctx context.Context, paymentData PPVEventPayment, payment *models.Payment) error {
	return nil
}

// RetrievePaymentStatus retrieves the payment made for an event
func (api api) RetrievePaymentStatus(ctx context.Context, eventID int64, user security.JWTUser) (models.PaymentSlice, error) {
	payments, err := models.Payments(
		qm.Where(models.PaymentColumns.PayPerViewEventID+" = ?", eventID),
		qm.Where(models.PaymentColumns.UserID+" = ?", user.ID),
	).All(ctx, api.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return payments, ErrInexistentPayment
		}
		return payments, err
	}
	return payments, nil
}

func (api api) RegisterSubscription(ctx context.Context, payment PPVEventPayment) (*models.Payment, error) {
	whereConditions := make([]qm.QueryMod, 0, 3)
	whereConditions = append(whereConditions, qm.Where(models.PaymentColumns.WalletAddress+" = ? AND "+models.PaymentColumns.PayPerViewEventID+" = ? AND "+models.PaymentColumns.Currency+" = ?", payment.WalletAddress, payment.EventID, payment.CryptoCurrency))
	if payment.User.ID != 0 {
		whereConditions = append(whereConditions, qm.And("user_id = ?", payment.User.ID))
	}
	if len(payment.PaymentID) > 0 {
		whereConditions = append(whereConditions, qm.And("currency_payment_id = ?", payment.PaymentID))
	}
	models.Payments(whereConditions...)
	return nil, nil
}

func (api api) RegisterUnsubscription(ctx context.Context, payment PPVEventPayment) (*models.Payment, error) {
	return nil, nil
}

func (api api) retrieveEventsFromDatabase(ctx context.Context, ppvQuery PPVEventQuery, user *security.JWTUser) ([]PPVEventAndSubscription, error) {
	subscription := make([]PPVEventAndSubscription, 0)
	results := make([]ppvEventAndSubscriptionResult, 0)
	queryBuilder := strings.Builder{}
	queryParams := make([]interface{}, 0, 4)
	queryBuilder.WriteString(`
	SELECT 
	pay_per_view_events.id AS "pay_per_view_events.id",
	pay_per_view_events.name AS "pay_per_view_events.name",
	pay_per_view_events.event_type AS "pay_per_view_events.event_type",
	pay_per_view_events.description AS "pay_per_view_events.description",
	pay_per_view_events.start AS "pay_per_view_events.start",
	pay_per_view_events.end AS "pay_per_view_events.end",
	pay_per_view_events.price_eth AS "pay_per_view_events.price_eth",
	pay_per_view_events.price_btc AS "pay_per_view_events.price_btc",
	pay_per_view_events.price_xmr AS "pay_per_view_events.price_xmr",
	pay_per_view_events.created_at AS "pay_per_view_events.createdAt",
	pay_per_view_events.updated_at AS "pay_per_view_events.updatedAt"
	`)
	paramN := 1
	if user != nil {
		queryParams = append(queryParams, user.ID)
		// LEFT OUTER JOIN payments ON (pay_per_view_events.id = payments.pay_per_view_event_id AND payments.user_id = $%d)
		queryBuilder.WriteString(`,
		payments.id AS "payments.id",
		payments.currency AS "payments.currency",
		payments.currency_payment_id AS "payments.currency_payment_id",
		payments.block_hash AS "payments.block_hash",
		payments.block_number_hex AS "payments.block_number_hex",
		payments.tx_hash AS "payments.tx_hash",
		payments.wallet_address AS "payments.wallet_address",
		payments.tx_number_hex AS "payments.tx_number_hex",
		payments.amount AS "payments.amount",
		payments.status AS "payments.status",
		payments.block_hash AS "payments.block_hash",
		payments.block_number_hex AS "payments.block_number_hex",
		payments.tx_hash AS "payments.tx_hash",
		payments.tx_number_hex AS "payments.tx_number_hex",
		payments.cancelled_block_hash AS "payments.cancelled_block_hash",
		payments.cancelled_block_number_hex AS "payments.cancelled_block_number_hex",
		payments.cancelled_tx_hash AS "payments.cancelled_tx_hash",
		payments.cancelled_tx_number_hex AS "payments.cancelled_tx_number_hex",
		payments.cancelled_at AS "payments.cancelled_at",
		payments.created_at AS "payments.created_at",
		payments.updated_at AS "payments.updated_at"
		FROM pay_per_view_events
	`)
		queryBuilder.WriteString("LEFT OUTER JOIN ")
		queryBuilder.WriteString(models.TableNames.Payments)
		queryBuilder.WriteString(" ON (")
		queryBuilder.WriteString(models.TableNames.PayPerViewEvents)
		queryBuilder.WriteString(".")
		queryBuilder.WriteString(models.PayPerViewEventColumns.ID)
		queryBuilder.WriteString(" = ")
		queryBuilder.WriteString(models.TableNames.Payments)
		queryBuilder.WriteString(".")
		queryBuilder.WriteString(models.PaymentColumns.PayPerViewEventID)
		queryBuilder.WriteString(" AND ")
		queryBuilder.WriteString(models.TableNames.Payments)
		queryBuilder.WriteString(".")
		queryBuilder.WriteString(models.PaymentColumns.UserID)
		queryBuilder.WriteString(" = $")
		queryBuilder.WriteString(strconv.Itoa(paramN))
		queryBuilder.WriteString(")\n")
		paramN++
	} else {
		queryBuilder.WriteString("FROM pay_per_view_events\n")
	}
	queryBuilder.WriteString("WHERE 1=1\n")
	if ppvQuery.ID != 0 {
		queryParams = append(queryParams, ppvQuery.ID)
		queryBuilder.WriteString(fmt.Sprintf("AND pay_per_view_events.id = $%d", paramN))
		paramN++
	}
	if ppvQuery.Query != "" {
		queryStr := strings.ReplaceAll(ppvQuery.Query, "%", "\\%") + "%"
		queryParams = append(queryParams, queryStr)
		queryBuilder.WriteString("AND (")
		queryBuilder.WriteString(models.TableNames.PayPerViewEvents)
		queryBuilder.WriteString(".")
		queryBuilder.WriteString(models.PayPerViewEventColumns.Name)
		queryBuilder.WriteString(fmt.Sprintf(" LIKE $%d OR ", paramN))
		queryBuilder.WriteString(models.TableNames.PayPerViewEvents)
		queryBuilder.WriteString(".")
		queryBuilder.WriteString(models.PayPerViewEventColumns.EventType)
		queryBuilder.WriteString(fmt.Sprintf(" LIKE $%d ", paramN))
		queryBuilder.WriteString(")\n")
		paramN++
	}
	if !ppvQuery.From.IsZero() {
		queryParams = append(queryParams, ppvQuery.From)
		queryBuilder.WriteString("AND ")
		queryBuilder.WriteString(models.TableNames.PayPerViewEvents)
		queryBuilder.WriteString(".")
		queryBuilder.WriteString(models.PayPerViewEventColumns.End)
		queryBuilder.WriteString(fmt.Sprintf(" >= $%d\n", paramN))
		paramN++
	}
	if !ppvQuery.To.IsZero() {
		queryParams = append(queryParams, ppvQuery.To)
		queryBuilder.WriteString("AND ")
		queryBuilder.WriteString(models.TableNames.PayPerViewEvents)
		queryBuilder.WriteString(".")
		queryBuilder.WriteString(models.PayPerViewEventColumns.End)
		queryBuilder.WriteString(fmt.Sprintf(" >= $%d\n", paramN))
		paramN++
	}
	queryBuilder.WriteString("ORDER BY pay_per_view_events.start, pay_per_view_events.id ")
	queryStr := queryBuilder.String()
	err := queries.Raw(queryStr, queryParams...).Bind(ctx, api.db, &results)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			api.logger.Printf("error running query in retrieveEventsFromDatabase: %v\n", err)
		}
		return subscription, err
	}
	var currentPPVEvent PPVEventAndSubscription
	if len(results) > 0 {
		for i := range results {
			event := results[i]
			if event.ID == currentPPVEvent.ID {
				if event.PID.Valid {
					currentPPVEvent.Payments = append(currentPPVEvent.Payments, fromNullabletoPayment(event))
				}
				subscription = append(subscription, currentPPVEvent)
			} else {
				if currentPPVEvent.ID != 0 {
					subscription = append(subscription, currentPPVEvent)
					currentPPVEvent = PPVEventAndSubscription{}
				}
				currentPPVEvent = PPVEventAndSubscription{}
				currentPPVEvent.ID = event.ID
				currentPPVEvent.Name = event.Name
				currentPPVEvent.Description = event.Description
				currentPPVEvent.EventType = event.EventType
				currentPPVEvent.Start = event.Start
				currentPPVEvent.End = event.End
				currentPPVEvent.PriceEth = event.PriceEth
				currentPPVEvent.PriceBTC = event.PriceBTC
				currentPPVEvent.PriceXMR = event.PriceXMR
				currentPPVEvent.CreatedAt = event.CreatedAt
				currentPPVEvent.UpdatedAt = event.UpdatedAt
				if event.PID.Valid {
					currentPPVEvent.Payments = append(currentPPVEvent.Payments, fromNullabletoPayment(event))
				}
			}
		}
		subscription = append(subscription, currentPPVEvent)
	}
	return subscription, nil
}

func fromNullabletoPayment(r ppvEventAndSubscriptionResult) models.Payment {
	return models.Payment{
		ID:                      r.PID.Int64,
		UserID:                  r.PaymentUserID,
		PayPerViewEventID:       r.PaymentPayPerViewEventID.Int64,
		Currency:                r.PaymentCurrency.String,
		CurrencyPaymentID:       r.PaymentCurrencyPaymentID,
		Amount:                  r.PaymentAmount,
		WalletAddress:           r.PaymentWalletAddress,
		Status:                  r.PaymentStatus.String,
		BlockHash:               r.PaymentBlockHash,
		BlockNumberHex:          r.PaymentBlockNumberHex,
		TXHash:                  r.PaymentTXHash,
		TXNumberHex:             r.PaymentTXNumberHex,
		CancelledBlockHash:      r.PaymentCancelledBlockHash,
		CancelledBlockNumberHex: r.PaymentCancelledBlockNumberHex,
		CancelledTXHash:         r.PaymentCancelledTXHash,
		CancelledTXNumberHex:    r.PaymentCancelledTXNumberHex,
		CancelledAt:             r.PaymentCancelledAt,
		CreatedAt:               r.PaymentCreatedAt,
		UpdatedAt:               r.PaymentUpdatedAt,
	}
}
