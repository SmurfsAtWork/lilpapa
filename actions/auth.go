package actions

import (
	"fmt"
	"strconv"
	"time"

	"github.com/SmurfsAtWork/lilpapa/app/models"
	"golang.org/x/crypto/bcrypt"
)

const (
	userSessionTokenTtlDays  = 60
	smurfSessionTokenTtlDays = 3600 // 10 years, basically non expiring lol
)

func (a *Actions) AuthenticateUser(sessionToken string) (models.User, error) {
	token, err := a.jwt.Decode(sessionToken, JwtSessionToken)
	if err != nil {
		return models.User{}, ErrInvalidSessionToken{}
	}

	if !token.Payload.Valid() {
		return models.User{}, ErrInvalidSessionToken{}
	}

	user, err := a.cache.GetAuthenticatedUser(sessionToken)
	if err != nil {
		user, err = a.app.GetUserByUsername(token.Payload.Username)
		if err != nil {
			return models.User{}, err
		}

		err = a.cache.SetAuthenticatedUser(sessionToken, user)
		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (a *Actions) AuthenticateSmurf(sessionToken string) (models.Smurf, error) {
	token, err := a.jwt.Decode(sessionToken, JwtSessionToken)
	if err != nil {
		return models.Smurf{}, err
	}

	if !token.Payload.Valid() {
		return models.Smurf{}, ErrInvalidSessionToken{}
	}

	smurf, err := a.cache.GetAuthenticatedSmurf(sessionToken)
	if err != nil {
		smurfId, _ := strconv.Atoi(token.Payload.Username)
		smurf, err = a.app.GetSmurfById(uint(smurfId))
		if err != nil {
			return models.Smurf{}, err
		}

		err = a.cache.SetAuthenticatedSmurf(sessionToken, smurf)
		if err != nil {
			return models.Smurf{}, err
		}
	}

	return smurf, nil
}

type TokenPayload struct {
	Username string `json:"username"`
}

func (t TokenPayload) Valid() bool {
	return t.Username != ""
}

type LoginUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserPayload struct {
	SessionToken string `json:"session_token"`
}

func (a *Actions) LoginUser(params LoginUserParams) (LoginUserPayload, error) {
	admin, err := a.app.GetUserByUsername(params.Username)
	if err != nil {
		return LoginUserPayload{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(params.Password))
	if err != nil {
		return LoginUserPayload{}, ErrInvalidLoginCredientials{}
	}

	sessionToken, err := a.jwt.Sign(TokenPayload{
		Username: params.Username,
	}, JwtSessionToken, time.Hour*24*userSessionTokenTtlDays)
	if err != nil {
		return LoginUserPayload{}, err
	}

	return LoginUserPayload{
		SessionToken: sessionToken,
	}, nil
}

type LoginSmurfParams struct {
	Id       uint   `json:"id"`
	Password string `json:"password"`
}

type LoginSmurfPayload struct {
	SessionToken string `json:"session_token"`
}

func (a *Actions) LoginSmurf(params LoginSmurfParams) (LoginSmurfPayload, error) {
	smurf, err := a.app.GetSmurfById(params.Id)
	if err != nil {
		return LoginSmurfPayload{}, err
	}

	if smurf.Password != params.Password {
		return LoginSmurfPayload{}, ErrInvalidLoginCredientials{}
	}

	sessionToken, err := a.jwt.Sign(TokenPayload{
		Username: fmt.Sprint(params.Id),
	}, JwtSessionToken, time.Hour*24*smurfSessionTokenTtlDays)
	if err != nil {
		return LoginSmurfPayload{}, err
	}

	return LoginSmurfPayload{
		SessionToken: sessionToken,
	}, nil
}

func (a *Actions) InvalidateAuthenticatedUser(sessionToken string) error {
	return a.cache.InvalidateAuthenticatedUser(sessionToken)
}

func (a *Actions) InvalidateAuthenticatedSmurf(sessionToken string) error {
	return a.cache.InvalidateAuthenticatedSmurf(sessionToken)
}
