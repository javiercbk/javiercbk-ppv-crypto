package security

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// UserNotFoundError is returned when a jwt token was not found in the request context
type UserNotFoundError string

func (e UserNotFoundError) Error() string {
	return string(e)
}

// MalformedUserError is returned when a user cannot be parsed from the JWT user
type MalformedUserError string

func (e MalformedUserError) Error() string {
	return string(e)
}

const (
	contextKey    = "jwtUser"
	userID        = "id"
	userFirstName = "firstName"
	userLastName  = "lastName"
	userIsAdmin   = "isAdmin"
	// ErrUserNotFound is returned when a jwt token was not found in the request context
	ErrUserNotFound UserNotFoundError = "user was not found in the request context"
	// ErrMalformedUser is returned when a user cannot be parsed from the JWT user
	ErrMalformedUser MalformedUserError = "user data is malformed"
)

// JWTUser is the data being encoded in the JWT token
type JWTUser struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	IsAdmin   string `json:"isAdmin"`
}

// JWTMiddlewareFactory creates a JWTMiddleware
func JWTMiddlewareFactory(jwtSecret string, optional bool) echo.MiddlewareFunc {
	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(jwtSecret),
		ContextKey: contextKey,
	})
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := jwtMiddleware(next)(c)
			if errors.Is(err, middleware.ErrJWTMissing) && optional {
				// if it failed to find the JWTToken, then continue
				// if and only if the user is optional
				return next(c)
			}
			return err
		}
	}
}

// JWTEncode encodes a user into a jwt.MapClaims
func JWTEncode(user JWTUser, d time.Duration) jwt.MapClaims {
	claims := jwt.MapClaims{}
	claims[userID] = user.ID
	claims[userFirstName] = user.FirstName
	claims[userLastName] = user.LastName
	claims[userIsAdmin] = user.IsAdmin
	// session lasts only 20 minutes
	claims["exp"] = time.Now().Add(d).Unix()
	return claims
}

// JWTDecode attempt to decode a user
func JWTDecode(c echo.Context, jwtUser *JWTUser) error {
	var err error
	var ok bool
	user := c.Get(contextKey).(*jwt.Token)
	if user == nil {
		err = ErrUserNotFound
	} else {
		claims := user.Claims.(jwt.MapClaims)
		if jwtUser.ID, ok = claims[userID].(int64); !ok {
			err = ErrMalformedUser
		}
		if jwtUser.FirstName, ok = claims[userFirstName].(string); !ok {
			err = ErrMalformedUser
		}
		if jwtUser.LastName, ok = claims[userFirstName].(string); !ok {
			err = ErrMalformedUser
		}
	}
	return err
}
