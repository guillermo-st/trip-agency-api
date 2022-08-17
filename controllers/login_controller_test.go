package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guillermo-st/trip-agency-api/models"
	"github.com/guillermo-st/trip-agency-api/repositories"
	"github.com/guillermo-st/trip-agency-api/services"
)

func getTestController(t testing.TB) (*LoginController, error) {
	t.Helper()
	initialUsr := models.User{
		Id:        1,
		FirstName: "Juan",
		LastName:  "Juanez",
		Email:     "juan_juanez@mail.com",
		Password:  "1234",
		RoleId:    1,
	}
	err := initialUsr.HashPassword()
	if err != nil {
		return nil, err
	}
	jwts, err := services.NewJsonWebTokenService()
	if err != nil {
		return nil, err
	}

	usrRepo := repositories.NewMockRepository(initialUsr)
	return NewLoginController(usrRepo, *jwts), nil
}

func TestLoginController(t *testing.T) {
	ctrl, _ := getTestController(t)
	gin.SetMode(gin.TestMode)

	t.Run("Valid login with existing user should result in StatusOK response", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		testBody := `{
    					"email": "juan_juanez@mail.com",
    					"password": "1234"
					}`

		ctx.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte(testBody)))
		ctrl.Login(ctx)

		if w.Code != http.StatusOK {
			t.Errorf("Unable to login with existing user")
		}
	})

	t.Run("Login with non-existing user should result in Not Found response", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		testBody := `{
    					"email": "mauro@mail.com",
    					"password": "1234"
					}`

		ctx.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte(testBody)))
		ctrl.Login(ctx)

		if w.Code != http.StatusNotFound {
			t.Errorf("Erroneous status code when login with non existing user %v", w.Code)
		}
	})

	t.Run("Login with existing user but wrong password should result in Unauthorized response", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		testBody := `{
    					"email": "juan_juanez@mail.com",
    					"password": "12"
					}`

		ctx.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte(testBody)))
		ctrl.Login(ctx)

		if w.Code != http.StatusUnauthorized {
			t.Errorf("Erroneous status code when login with wrong password %v", w.Code)
		}
	})

	t.Run("Login with erroneous body should result in Bad Request response", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		testBody := `{
					}`

		ctx.Request, _ = http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte(testBody)))
		ctrl.Login(ctx)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Erroneous status code when login with wrong body %v", w.Code)
		}
	})
}
