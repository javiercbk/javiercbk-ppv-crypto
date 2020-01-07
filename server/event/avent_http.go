package event

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"time"

	"github.com/javiercbk/ppv-crypto/server/http/response"
	"github.com/javiercbk/ppv-crypto/server/http/security"
	"github.com/javiercbk/ppv-crypto/server/models"
	"github.com/labstack/echo/v4"
)

const (
	eventIDParamKey = "eventID"
)

func parseEventID(c echo.Context) (int64, error) {
	eventIDStr := c.Param(eventIDParamKey)
	eventID, err := strconv.ParseInt(eventIDStr, 10, 64)
	return eventID, err
}

// ProspectEvent contains the information to create an pay per view event
type ProspectEvent struct {
	Name            string    `json:"name" validate:"required,gt=0,lte=200"`
	Description     string    `json:"description" validate:"required,gt=0,lte=1500"`
	EventType       string    `json:"eventType" validate:"required,gt=0,lte=200"`
	Start           time.Time `json:"start" validate:"required"`
	End             time.Time `json:"end" validate:"required"`
	PriceETH        int64     `json:"priceETH" validate:"required"`
	PriceBTC        int64     `json:"priceBTC" validate:"required"`
	PriceXMR        int64     `json:"priceXMR" validate:"required"`
	EthContractAddr string    `json:"ethContractAddr" validate:"required,gt=0"`
}

// Handler is a group of handlers within a route.
type Handler struct {
	logger     *log.Logger
	db         *sql.DB
	apiFactory APIFactory
}

// NewHandler creates a handler for the game route
func NewHandler(logger *log.Logger, db *sql.DB) Handler {
	return Handler{
		logger:     logger,
		db:         db,
		apiFactory: NewAPI,
	}
}

// Routes initializes all the routes with their http handlers
func (h Handler) Routes(e *echo.Group, jwtMiddleware echo.MiddlewareFunc, jwtOptionalMiddleware echo.MiddlewareFunc) {
	e.GET("", h.retrieveEventList, jwtOptionalMiddleware)
	e.POST("", h.createEvent, jwtMiddleware)
	e.GET("/:eventID", h.retrieveEvent, jwtOptionalMiddleware)
	e.POST("/:eventID/payment", h.processPayment, jwtMiddleware)
	e.GET("/:eventID/payment", h.retrievePayments, jwtMiddleware)
}

// retrieveEventList retrieves the list of events
func (h Handler) retrieveEventList(c echo.Context) error {
	ctx := c.Request().Context()
	jwtUser := &security.JWTUser{}
	err := security.JWTDecode(c, jwtUser)
	if err != nil {
		if !errors.Is(err, security.ErrUserNotFound) {
			h.logger.Printf("could not parse user%v\n", err)
			return response.NewUnauthorizedErrorResponse(c)
		}
		jwtUser = nil
		err = nil
	}
	ppvQuery := PPVEventQuery{}
	err = c.Bind(&ppvQuery)
	if err != nil {
		h.logger.Printf("could not bind request data%v\n", err)
		return response.NewBadRequestResponse(c, err.Error())
	}
	if err = c.Validate(ppvQuery); err != nil {
		h.logger.Printf("validation error %v\n", err)
		return response.NewBadRequestResponse(c, err.Error())
	}
	api := h.apiFactory(h.logger, h.db)
	events, err := api.RetrieveEvents(ctx, ppvQuery, jwtUser)
	if err != nil {
		h.logger.Printf("error retrieving events %v\n", err)
		return response.NewInternalErrorResponse(c, err.Error())
	}
	return response.NewSuccessResponse(c, events)
}

func (h Handler) createEvent(c echo.Context) error {
	ctx := c.Request().Context()
	jwtUser := &security.JWTUser{}
	err := security.JWTDecode(c, jwtUser)
	if err != nil {
		return response.NewUnauthorizedErrorResponse(c)
	}
	prospectEvent := ProspectEvent{}
	err = c.Bind(&prospectEvent)
	if err != nil {
		h.logger.Printf("could not bind request data%v\n", err)
		return response.NewBadRequestResponse(c, err.Error())
	}
	if err = c.Validate(prospectEvent); err != nil {
		h.logger.Printf("validation error %v\n", err)
		return response.NewBadRequestResponse(c, err.Error())
	}
	api := h.apiFactory(h.logger, h.db)
	ppvEvent := &models.PayPerViewEvent{}
	err = api.CreateEvent(ctx, ppvEvent)
	if err != nil {
		h.logger.Printf("error creating event %v\n", err)
		return response.NewInternalErrorResponse(c, err.Error())
	}
	return response.NewSuccessResponse(c, ppvEvent)
}

// retrieveEvent retrieves an event
func (h Handler) retrieveEvent(c echo.Context) error {
	ctx := c.Request().Context()
	eventID, err := parseEventID(c)
	if err != nil {
		return response.NewNotFoundResponse(c)
	}
	jwtUser := &security.JWTUser{}
	err = security.JWTDecode(c, jwtUser)
	if err != nil {
		if !errors.Is(err, security.ErrUserNotFound) {
			h.logger.Printf("could not parse user%v\n", err)
			return response.NewUnauthorizedErrorResponse(c)
		}
		jwtUser = nil
		err = nil
	}
	api := h.apiFactory(h.logger, h.db)
	event, err := api.RetrieveEvent(ctx, eventID, jwtUser)
	if err != nil {
		if errors.Is(err, ErrInexistentEvent) {
			return response.NewNotFoundResponse(c)
		}
		return response.NewInternalErrorResponse(c, err.Error())
	}
	return response.NewSuccessResponse(c, event)
}

// processPayment creates an intent to process a payment
func (h Handler) processPayment(c echo.Context) error {
	// ctx := c.Request().Context()
	eventID, err := parseEventID(c)
	if err != nil {
		return response.NewNotFoundResponse(c)
	}
	jwtUser := &security.JWTUser{}
	err = security.JWTDecode(c, jwtUser)
	if err != nil {
		return response.NewUnauthorizedErrorResponse(c)
	}
	h.logger.Printf("about to process payment for event %d\n", eventID)
	return response.NewSuccessResponse(c, nil)
}

// retrievePayments retrieves the payments done to an event by a user
func (h Handler) retrievePayments(c echo.Context) error {
	ctx := c.Request().Context()
	eventID, err := parseEventID(c)
	if err != nil {
		return response.NewNotFoundResponse(c)
	}
	jwtUser := security.JWTUser{}
	err = security.JWTDecode(c, &jwtUser)
	if err != nil {
		return response.NewUnauthorizedErrorResponse(c)
	}
	api := h.apiFactory(h.logger, h.db)
	payments, err := api.RetrievePaymentStatus(ctx, eventID, jwtUser)
	return response.NewSuccessResponse(c, payments)
}
