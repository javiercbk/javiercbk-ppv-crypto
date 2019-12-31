package event

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/javiercbk/ppv-crypto/server/http/security"
	"github.com/javiercbk/ppv-crypto/server/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

// ErrDuplicatedUser is returned when attempting to create a duplicated user
type ErrDuplicatedUser string

func (e ErrDuplicatedUser) Error() string {
	return string(e)
}

// NameTooLongErr is returned when attempting to store a first or last name that is too long
type NameTooLongErr string

func (e NameTooLongErr) Error() string {
	return string(e)
}

// EmailTooLongErr is returned when attempting to store an email that is too long
type EmailTooLongErr string

func (e EmailTooLongErr) Error() string {
	return string(e)
}

// BadCredentialsErr is returned when attempting to login with invalid credentials
type BadCredentialsErr string

func (e BadCredentialsErr) Error() string {
	return string(e)
}

const (
	nameConstraint        = "user_name_cnst"
	emailConstraint       = "user_email_cnst"
	emailUniqueConstraint = "users_email_idx"
	// ErrNameTooLong is returned when attempting to store a first or last name that is too long
	ErrNameTooLong NameTooLongErr = "either first or last name is too long"
	// ErrEmailTooLong is returned when attempting to store an email that is too long
	ErrEmailTooLong EmailTooLongErr = "email is too long"
	// ErrBadCredentials is returned when attempting to login with invalid credentials
	ErrBadCredentials BadCredentialsErr = "bad credentials"
)

// AuthCredentials has all the data necesary to authenticate an admin
type AuthCredentials struct {
	Email    string `json:"email" validate:"required,gt=0,lte=256"`
	Password string `json:"password" validate:"required,gt=0"`
}

// TokenResponse contains a jwt token
type TokenResponse struct {
	User  security.JWTUser `json:"user"`
	Token string           `json:"token"`
}

// VisibleAdmin is the public data of an admin
type VisibleUser struct {
	ID        int64      `json:"id"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Email     string     `json:"email"`
	IsAdmin   bool       `json:"isAdmin"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// APIFactory is a function capable of creating an Auth API
type APIFactory func(logger *log.Logger, db *sql.DB, jwtSecret string) API

// API is authentication API interface
type API interface {
	AuthenticateUser(ctx context.Context, credentials AuthCredentials) (TokenResponse, error)
}

type api struct {
	logger    *log.Logger
	db        *sql.DB
	jwtSecret string
}

// NewAPI creates a new authentication API
func NewAPI(logger *log.Logger, db *sql.DB) API {
	return api{
		logger: logger,
		db:     db,
	}
}

func (api api) AuthenticateUser(ctx context.Context, credentials AuthCredentials) (TokenResponse, error) {
	tokenResponse := TokenResponse{}
	user, err := models.Users(
		qm.Where("email = ?", credentials.Email),
	).One(ctx, api.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// if the admin does not exist, in this API method return unvalid credentials
			return tokenResponse, ErrBadCredentials
		}
		api.logger.Printf("error retrieving admin for update: %v\n", err)
		return tokenResponse, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return tokenResponse, ErrBadCredentials
	}
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = security.JWTEncode(security.JWTUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		IsAdmin:   user.IsAdmin,
	}, time.Minute*20)

	t, err := token.SignedString([]byte(api.jwtSecret))
	if err != nil {
		api.logger.Printf("error signing token %v\n", err)
		return tokenResponse, errors.New("error creating token")
	}
	tokenResponse.Token = t
	tokenResponse.User.ID = user.ID
	tokenResponse.User.FirstName = user.FirstName
	tokenResponse.User.LastName = user.LastName
	tokenResponse.User.IsAdmin = user.IsAdmin
	return tokenResponse, nil

}
