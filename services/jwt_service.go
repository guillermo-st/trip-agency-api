package services

import (
	"encoding/json"
	"errors"
	"time"

	jwt "github.com/cristalhq/jwt/v4"
	"github.com/guillermo-st/trip-agency-api/models"
)

type TokenClaims struct {
	IssuedAt  time.Time `json:"issued_at"`
	ExpiresAt time.Time `json:"expires_at"`
	UserId    uint      `json:"user_id"`
	IsAdmin   bool      `json:"is_admin"`
}

type JsonWebTokenService struct {
	secret   []byte
	builder  jwt.Builder
	verifier jwt.Verifier
}

func NewJsonWebTokenService() (*JsonWebTokenService, error) {
	secret := []byte("thisShouldBeAnEnvVariable")
	signer, err := jwt.NewSignerHS(jwt.HS256, secret)
	if err != nil {
		return nil, err
	}

	verifier, err := jwt.NewVerifierHS(jwt.HS256, secret)
	if err != nil {
		return nil, err
	}

	builder := jwt.NewBuilder(signer)

	return &JsonWebTokenService{secret, *builder, verifier}, nil
}

func (jwts *JsonWebTokenService) GenerateTokenWithUserClaims(usr models.User) (string, error) {
	claims := &TokenClaims{
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 8),
		UserId:    usr.Id,
		IsAdmin:   (usr.RoleId == 1),
	}

	token, err := jwts.builder.Build(claims)
	return token.String(), err
}

func (jwts *JsonWebTokenService) ValidateAndParseToken(rawToken string) (*TokenClaims, error) {
	token, err := jwt.Parse([]byte(rawToken), jwts.verifier)
	if err != nil {
		return nil, errors.New("Invalid Token. Please re-login and try again.")
	}

	var claims TokenClaims
	err = json.Unmarshal(token.Claims(), &claims)
	if err != nil {
		return nil, errors.New("Invalid Token. Please re-login and try again.")
	}

	if claims.ExpiresAt.Before(time.Now()) {
		return nil, errors.New("Your access Token has expired. Please re-login and try again.")
	}

	return &claims, nil
}
