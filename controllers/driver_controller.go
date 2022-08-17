package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guillermo-st/trip-agency-api/models"
	"github.com/guillermo-st/trip-agency-api/repositories"
)

type DriverController struct {
	repo repositories.DriverRepository
}

func NewDriverController(repo repositories.DriverRepository) *DriverController {
	return &DriverController{repo}
}

func (ctrl *DriverController) GetDrivers(ctx *gin.Context) {
	type request struct {
		PageNum  uint `json:"page_num" binding:"required"`
		PageSize uint `json:"page_size" binding:"required"`
	}
	var req request

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}

	drivers, err := ctrl.repo.GetDrivers(req.PageSize, req.PageNum)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not retrieve Driver Data. Please submit a support ticket.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"drivers": drivers,
	})
}

func (ctrl *DriverController) GetDriversByStatus(ctx *gin.Context) {
	type request struct {
		PageNum  uint `json:"page_num" binding:"required"`
		PageSize uint `json:"page_size" binding:"required"`
		IsOnTrip bool `json:"is_on_trip" binding:"required"`
	}
	var req request

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}

	drivers, err := ctrl.repo.GetDriversByStatus(req.PageSize, req.PageNum, req.IsOnTrip)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not retrieve Driver Data. Please submit a support ticket.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"drivers": drivers,
	})
}

func (ctrl *DriverController) AddDriver(ctx *gin.Context) {
	type request struct {
		Driver models.User `json:"driver"`
	}
	var req request

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}

	if err := req.Driver.HashPassword(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not insert new driver data. Please submit a support ticket.",
		})
		return
	}

	err := ctrl.repo.AddDriver(req.Driver)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": "Could not insert new driver data. Please submit a support ticket.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"error":   false,
		"message": "New driver succesfully registered!",
	})
}
