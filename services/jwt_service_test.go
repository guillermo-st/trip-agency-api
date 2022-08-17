package services

import (
	"testing"

	"github.com/cristalhq/jwt/v4"
	"github.com/guillermo-st/trip-agency-api/models"
)

func getTestUser(t testing.TB) models.User {
	t.Helper()
	return models.User{
		Id:        1,
		FirstName: "Bob",
		LastName:  "Bobbinson",
		Email:     "bob_bobbinson@mail.com",
		Password:  "hunter2",
		RoleId:    1,
	}
}

func TestJWTService(t *testing.T) {
	jwts, _ := NewJsonWebTokenService()

	t.Run("Generating token from user should result in a valid JWT token", func(t *testing.T) {
		rawToken, err := jwts.GenerateTokenWithUserClaims(getTestUser(t))
		if err != nil {
			t.Errorf("Couldn't generate token from User.")
		}
		token, err := jwt.Parse([]byte(rawToken), jwts.verifier)
		if err != nil {
			t.Errorf("Couldn't parse the User token.")
		}
		verifyErr := jwts.verifier.Verify(token)
		if verifyErr != nil {
			t.Errorf("Invalid Token generated from user")
		}
	})

	t.Run("Validating a Token generated from User should not result in error", func(t *testing.T) {
		usr := getTestUser(t)
		rawToken, err := jwts.GenerateTokenWithUserClaims(usr)
		if err != nil {
			t.Errorf("Couldn't generate token from User.")
		}

		_, err = jwts.ValidateAndParseToken(rawToken)
		if err != nil {
			t.Errorf("Couldn't validate the user token.")
		}
	})

	t.Run("Validating invalid JWT  Token should result in error", func(t *testing.T) {
		_, err := jwts.ValidateAndParseToken("invalidToken")
		if err == nil {
			t.Errorf("Invalid token shouldn't be succesfully validated")
		}
	})

	t.Run("Validating a Token generated from User should result in equivalent token claims", func(t *testing.T) {
		usr := getTestUser(t)
		rawToken, err := jwts.GenerateTokenWithUserClaims(usr)
		if err != nil {
			t.Errorf("Couldn't generate token from User.")
		}

		claims, err := jwts.ValidateAndParseToken(rawToken)
		if err != nil {
			t.Errorf("Couldn't validate the user token.")
		}

		if claims.UserId != usr.Id || claims.IsAdmin != (usr.RoleId == models.ADMIN_ROLE_ID) {
			t.Errorf("Incompatible claims inside generated JWT.")
		}
	})
}
