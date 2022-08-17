package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guillermo-st/trip-agency-api/repositories"
)

type TripController struct {
	repo repositories.TripRepository
}

func NewTripController(repo repositories.TripRepository) *TripController {
	return &TripController{repo}
}

func (ctrl *TripController) StartNewTripForDriver(ctx *gin.Context) {
	driverId, _ := ctx.Get("user_id")

	err := ctrl.repo.StartNewTripForDriver(driverId.(uint))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "New trip for driver succesfully registered!",
	})
}

func (ctrl *TripController) FinishTripForDriver(ctx *gin.Context) {
	driverId, _ := ctx.Get("user_id")
	err := ctrl.repo.FinishTripForDriver(driverId.(uint))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "Succesfully finished Trip for driver!",
	})
}
