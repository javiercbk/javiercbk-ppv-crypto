package auth

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/javiercbk/ppv-crypto/server/http/response"
	"github.com/javiercbk/ppv-crypto/server/http/security"
	"github.com/labstack/echo/v4"
)

// Handler is a group of handlers within a route.
type Handler struct {
	logger     *log.Logger
	db         *sql.DB
	apiFactory APIFactory
	jwtSecret  string
}

// NewHandler creates a handler for the game route
func NewHandler(logger *log.Logger, db *sql.DB, jwtSecret string) Handler {
	return Handler{
		logger:     logger,
		db:         db,
		apiFactory: NewAPI,
		jwtSecret:  jwtSecret,
	}
}

// Routes initializes all the routes with their http handlers
func (h Handler) Routes(e *echo.Group, jwtMiddleware echo.MiddlewareFunc) {
	e.POST("", h.authenticateUser)
	e.GET("/current", h.retrieveCurrentUserInfo, jwtMiddleware)
}

// retrieveEventList retrieves the list of events
func (h Handler) authenticateUser(c echo.Context) error {
	ctx := c.Request().Context()
	credentials := Credentials{}
	err := c.Bind(&credentials)
	if err != nil {
		h.logger.Printf("could not bind request data%v\n", err)
		return response.NewBadRequestResponse(c, err.Error())
	}
	if err = c.Validate(credentials); err != nil {
		h.logger.Printf("validation error %v\n", err)
		return response.NewBadRequestResponse(c, err.Error())
	}
	api := h.apiFactory(h.logger, h.db, h.jwtSecret)
	events, err := api.AuthenticateUser(ctx, credentials)
	if err != nil {
		h.logger.Printf("error authentication user %v\n", err)
		if errors.Is(err, ErrBadCredentials) {
			return response.NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		}
		return response.NewInternalErrorResponse(c, err.Error())
	}
	return response.NewSuccessResponse(c, events)
}

func (h Handler) retrieveCurrentUserInfo(c echo.Context) error {
	ctx := c.Request().Context()
	jwtUser := &security.JWTUser{}
	err := security.JWTDecode(c, jwtUser)
	if err != nil {
		return response.NewUnauthorizedErrorResponse(c)
	}
	api := h.apiFactory(h.logger, h.db, h.jwtSecret)
	visibleUser := VisibleUser{}
	err = api.UserInfo(ctx, jwtUser.ID, &visibleUser)
	if err != nil {
		if errors.Is(err, ErrUserNotExist) {
			return response.NewUnauthorizedErrorResponse(c)
		}
		return response.NewInternalErrorResponse(c, err.Error())
	}
	return response.NewSuccessResponse(c, visibleUser)
}
