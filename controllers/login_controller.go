package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guillermo-st/trip-agency-api/repositories"
	"github.com/guillermo-st/trip-agency-api/services"
)

type LoginController struct {
	repo     repositories.UserRepository
	authServ services.JsonWebTokenService
}

func NewLoginController(repo repositories.UserRepository, authServ services.JsonWebTokenService) *LoginController {
	return &LoginController{repo, authServ}
}

func (ctrl *LoginController) Login(ctx *gin.Context) {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req request

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}

	user, err := ctrl.repo.GetUserByEmail(req.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   true,
			"message": "User does not exist.",
		})
		return
	}

	err = user.ValidatePassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "Invalid Password, please try again",
		})
		return
	}

	token, err := ctrl.authServ.GenerateTokenWithUserClaims(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Unable to login user due to server error. Please submit a support ticket.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": token,
	})
}
