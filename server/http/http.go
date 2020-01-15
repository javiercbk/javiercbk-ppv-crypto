package http

import (
	"context"
	"crypto/tls"
	"database/sql"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/javiercbk/ppv-crypto/server/auth"
	"github.com/javiercbk/ppv-crypto/server/event"
	"github.com/javiercbk/ppv-crypto/server/http/response"
	"github.com/javiercbk/ppv-crypto/server/http/security"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gommonLog "github.com/labstack/gommon/log"

	// imports the postgres sql driver
	_ "github.com/lib/pq"
)

// BodyLimit is the upper limit for the http request body size.
const BodyLimit = "12M"

// Config contains all the configurations to initialize an http server
type Config struct {
	Address   string
	JWTSecret string
}

type customValidator struct {
	validator *validator.Validate
}

func (cv *customValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	if errors.Is(err, echo.ErrNotFound) {
		response.NewNotFoundResponse(c)
	} else {
		if errors.Is(err, middleware.ErrJWTMissing) {
			response.NewUnauthorizedErrorResponse(c)
		} else if echoErr, ok := err.(*echo.HTTPError); ok {
			response.NewErrorResponseWithCode(c, echoErr.Code)
		} else {
			response.NewResponseFromError(c, err)
		}
	}
}

// Serve http connections
func Serve(cnf Config, logger *log.Logger, db *sql.DB) error {
	router := echo.New()
	router.HTTPErrorHandler = customHTTPErrorHandler
	router.Validator = &customValidator{validator: validator.New()}
	router.Logger.SetLevel(gommonLog.INFO)
	router.Use(middleware.Recover())
	router.Use(middleware.Secure())
	// set a body limit of 12 megabites
	router.Use(middleware.BodyLimit(BodyLimit))
	router.Use(middleware.Gzip())
	initRoutes(router, cnf.JWTSecret, logger, db)
	srv := newServer(router, cnf.Address)
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			router.Logger.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be catched, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	router.Logger.Printf("Shutdown Server ...\n")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		router.Logger.Fatal("Server Shutdown:", err)
	}
	<-ctx.Done()
	router.Logger.Printf("timeout of 1 seconds.\n")
	router.Logger.Printf("Server exiting\n")
	return nil
}

func initRoutes(router *echo.Echo, jwtSecret string, logger *log.Logger, db *sql.DB) {
	jwtOptionalMiddleware := security.JWTMiddlewareFactory(jwtSecret, true)
	jwtMiddleware := security.JWTMiddlewareFactory(jwtSecret, false)
	apiRouter := router.Group("/api/v1")
	{
		eventRouter := apiRouter.Group("/events")
		eventHandler := event.NewHandler(logger, db)
		eventHandler.Routes(eventRouter, jwtMiddleware, jwtOptionalMiddleware)
	}
	{
		authRouter := apiRouter.Group("/auth")
		authHandler := auth.NewHandler(logger, db, jwtSecret)
		authHandler.Routes(authRouter, jwtMiddleware)
	}
	// adminsRouter := apiRouter.Group("/admins")
	// adminsRouter.Use(jwtMiddleware)
	// {
	// 	adminHandler := admin.NewHandler(logger, db)
	// 	adminHandler.Routes(adminsRouter)
	// }
}

func newServer(handler http.Handler, address string) *http.Server {
	// I used to follow the recommendations on https://blog.cloudflare.com/exposing-go-on-the-internet/
	// but I'll have to review them because most of the issues mentioned there seem
	// to be already solved.
	tlsConfig := &tls.Config{}

	return &http.Server{
		Addr: address,
		// we allow 60 seconds of read timeout for long polling
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		Handler:      handler,
	}
}
