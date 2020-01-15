package security

import (
	"errors"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/javiercbk/ppv-crypto/server/models"
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

type permission string

const (
	contextKey      = "jwtUser"
	userID          = "id"
	userFirstName   = "firstName"
	userLastName    = "lastName"
	userPermissions = "permissions"
	// ErrUserNotFound is returned when a jwt token was not found in the request context
	ErrUserNotFound UserNotFoundError = "user was not found in the request context"
	// ErrMalformedUser is returned when a user cannot be parsed from the JWT user
	ErrMalformedUser MalformedUserError = "user data is malformed"
	// Read is a permission that allows the user to read a resource
	Read permission = "read"
	// Write is a permission that allows the user to write a resource
	Write permission = "write"
)

// PermissionMap is the user's permission map
type PermissionMap map[string][]permission

// JWTUser is the data being encoded in the JWT token
type JWTUser struct {
	ID          int64         `json:"id"`
	FirstName   string        `json:"firstName"`
	LastName    string        `json:"lastName"`
	Permissions PermissionMap `json:"permissions"`
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
	claims[userID] = strconv.FormatInt(user.ID, 10)
	claims[userFirstName] = user.FirstName
	claims[userLastName] = user.LastName
	claims[userPermissions] = user.Permissions
	// session lasts only 20 minutes
	claims["exp"] = time.Now().Add(d).Unix()
	return claims
}

// JWTDecode attempt to decode a user
func JWTDecode(c echo.Context, jwtUser *JWTUser) error {
	var err error
	user, ok := c.Get(contextKey).(*jwt.Token)
	if !ok {
		err = ErrUserNotFound
	} else {
		claims := user.Claims.(jwt.MapClaims)
		var idStr string
		if idStr, ok = claims[userID].(string); !ok {
			err = ErrMalformedUser
		}
		jwtUser.ID, err = strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			err = ErrMalformedUser
		}
		if jwtUser.FirstName, ok = claims[userFirstName].(string); !ok {
			err = ErrMalformedUser
		}
		if jwtUser.LastName, ok = claims[userFirstName].(string); !ok {
			err = ErrMalformedUser
		}
		var permissionMap map[string]interface{}
		if permissionMap, ok = claims[userPermissions].(map[string]interface{}); !ok {
			err = ErrMalformedUser
		}
		jwtUser.Permissions = make(PermissionMap)
		for key, val := range permissionMap {
			arr := val.([]interface{})
			jwtUser.Permissions[key] = make([]permission, len(arr))
			for i := range arr {
				permStr := arr[i].(string)
				if permStr == string(Write) {
					jwtUser.Permissions[key][i] = Write
				} else {
					jwtUser.Permissions[key][i] = Read
				}
			}
		}
	}
	return err
}

// CanReadResouceMiddleware returns a middleware that validates if a user can read a resource
func CanReadResouceMiddleware(jwtUser JWTUser, resouce string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if jwtUser.ID == 0 {
				return echo.ErrUnauthorized
			}
			if _, ok := jwtUser.Permissions[resouce]; ok {
				// if it has any permission, then it can read the resource
				return next(c)
			}
			return echo.ErrForbidden
		}
	}
}

// CanWriteResouceMiddleware returns a middleware that validates if a user can read/write a resource
func CanWriteResouceMiddleware(jwtUser JWTUser, resouce string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if jwtUser.ID == 0 {
				return echo.ErrUnauthorized
			}
			if permissions, ok := jwtUser.Permissions[resouce]; ok {
				// if it has any permission, then it can read the resource
				canWrite := false
				for i := range permissions {
					if permissions[i] == Write {
						canWrite = true
						break
					}
				}
				if canWrite {
					return next(c)
				}
			}
			return echo.ErrForbidden
		}
	}
}

// ToPermissionsMap creates a permissionsMap out of a PermissionsSlice
func ToPermissionsMap(permissions models.PermissionsUserSlice) PermissionMap {
	permMap := make(PermissionMap)
	for i := range permissions {
		prm := permissions[i]
		if permMap[prm.Resource] == nil {
			permMap[prm.Resource] = make([]permission, 1)
			permMap[prm.Resource][0] = parsePermission(prm.Access)
		} else {
			permMap[prm.Resource] = append(permMap[prm.Resource], parsePermission(prm.Access))
		}
	}
	return permMap
}

func parsePermission(access string) permission {
	if access == string(Write) {
		return Write
	}
	return Read
}
