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
	"github.com/javiercbk/ppv-crypto/server/http/security"
	"github.com/javiercbk/ppv-crypto/server/models"
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
type APIFactory func(logger *log.Logger, db *sql.DB) API

// PPVEventAndSubscription an event containing the user payments
type PPVEventAndSubscription struct {
	models.PayPerViewEvent
	Payments []models.Payment `json:"payments"`
}

// ppvEventAndSubscriptionResults is used to bind the query results from a query
type ppvEventAndSubscriptionResults struct {
	models.PayPerViewEvent `boil:",bind"`
	Payment                models.Payment `boil:",bind"`
	// models.PayPerViewEvent  `boil:"pay_per_view_events,bind"`
	// PaymentID               int64       `boil:"p.id"`
	// UserID                  int64       `boil:"p.user_id"`
	// PayPerViewEventID       int64       `boil:"p.pay_per_view_event_id"`
	// Currency                string      `boil:"p.currency"`
	// CurrencyPaymentID       null.String `boil:"p.currency_payment_id"`
	// Amount                  null.Int64  `boil:"p.amount"`
	// WalletAddress           null.String `boil:"p.wallet_address"`
	// Status                  string      `boil:"p.status"`
	// BlockHash               null.String `boil:"p.block_hash"`
	// BlockNumberHex          null.String `boil:"p.block_number_hex"`
	// TXHash                  null.String `boil:"p.tx_hash"`
	// TXNumberHex             null.String `boil:"p.tx_number_hex"`
	// CancelledBlockHash      null.String `boil:"p.cancelled_block_hash"`
	// CancelledBlockNumberHex null.String `boil:"p.cancelled_block_number_hex"`
	// CancelledTXHash         null.String `boil:"p.cancelled_tx_hash"`
	// CancelledTXNumberHex    null.String `boil:"p.cancelled_tx_number_hex"`
	// CancelledAt             null.Time   `boil:"p.cancelled_at"`
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
	logger *log.Logger
	db     *sql.DB
}

// NewAPI creates a new authentication API
func NewAPI(logger *log.Logger, db *sql.DB) API {
	return api{
		logger: logger,
		db:     db,
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
	err := ppvEvent.Insert(ctx, api.db, boil.Whitelist(
		models.PayPerViewEventColumns.Name,
		models.PayPerViewEventColumns.Description,
		models.PayPerViewEventColumns.EventType,
		models.PayPerViewEventColumns.Start,
		models.PayPerViewEventColumns.End,
		models.PayPerViewEventColumns.PriceEth,
		models.PayPerViewEventColumns.PriceBTC,
		models.PayPerViewEventColumns.PriceXMR,
		models.PayPerViewEventColumns.EthContractAddr,
		models.PayPerViewEventColumns.CreatedAt,
	))
	if err != nil {
		api.logger.Printf("error inserting event: %v\n", err)
	}
	return err
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
	results := make([]ppvEventAndSubscriptionResults, 0)
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
	pay_per_view_events.price_ETH AS "pay_per_view_events.price_ETH",
	pay_per_view_events.price_BTC AS "pay_per_view_events.price_BTC",
	pay_per_view_events.price_XMR AS "pay_per_view_events.price_XMR",
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
	api.logger.Println(queryStr)
	err := queries.Raw(queryStr, queryParams...).Bind(ctx, api.db, &results)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			api.logger.Printf("error running query in retrieveEventsFromDatabase: %v\n", err)
		}
		return subscription, err
	}
	var currentPPVEvent PPVEventAndSubscription
	for i := range results {
		event := results[i]
		if event.ID == currentPPVEvent.ID {
			currentPPVEvent.Payments = append(currentPPVEvent.Payments, event.Payment)
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
			if event.Payment.ID != 0 {
				currentPPVEvent.Payments = append(currentPPVEvent.Payments, event.Payment)
			}
		}
	}
	return subscription, nil
}
