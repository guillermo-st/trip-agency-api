package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/guillermo-st/trip-agency-api/controllers"
	"github.com/guillermo-st/trip-agency-api/middleware"
	"github.com/guillermo-st/trip-agency-api/repositories"
	"github.com/guillermo-st/trip-agency-api/services"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	loginCtrl  *controllers.LoginController
	tripCtrl   *controllers.TripController
	driverCtrl *controllers.DriverController
	router     *gin.Engine
}

func NewServer(userRepo repositories.UserRepository, tripRepo repositories.TripRepository, driverRepo repositories.DriverRepository) (*Server, error) {
	tripCtrl := controllers.NewTripController(tripRepo)
	driverCtrl := controllers.NewDriverController(driverRepo)
	authServ, err := services.NewJsonWebTokenService()
	if err != nil {
		return nil, err
	}
	loginCtrl := controllers.NewLoginController(userRepo, *authServ)
	srv := &Server{loginCtrl, tripCtrl, driverCtrl, gin.Default()}
	srv.initRoutes()
	return srv, nil
}

func DefaultServer(db *sqlx.DB) (*Server, error) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}

	tripRepo := repositories.NewDBTripRepository(db)
	tripCtrl := controllers.NewTripController(tripRepo)

	driverRepo := repositories.NewDBDriverRepository(db)
	driverCtrl := controllers.NewDriverController(driverRepo)

	userRepo := repositories.NewDBUserRepository(db)

	authServ, err := services.NewJsonWebTokenService()
	if err != nil {
		panic(err)
	}
	loginCtrl := controllers.NewLoginController(userRepo, *authServ)

	srv := &Server{loginCtrl, tripCtrl, driverCtrl, gin.Default()}
	srv.initRoutes()
	return srv, nil
}

func (srv *Server) initRoutes() {
	srv.router.GET("/drivers", middleware.AuthorizeJWT(true), srv.driverCtrl.GetDrivers)
	srv.router.GET("/drivers-by-status", middleware.AuthorizeJWT(true), srv.driverCtrl.GetDriversByStatus)
	srv.router.POST("/drivers", middleware.AuthorizeJWT(true), srv.driverCtrl.AddDriver)

	tripRoutes := srv.router.Group("/trips", middleware.AuthorizeJWT(false))
	{
		tripRoutes.POST("/start", srv.tripCtrl.StartNewTripForDriver)
		tripRoutes.POST("/finish", srv.tripCtrl.FinishTripForDriver)
	}

	srv.router.POST("/trips")

	srv.router.POST("/login", srv.loginCtrl.Login)
}

func (srv *Server) Run(port uint) error {
	portStr := fmt.Sprintf(":%v", port)
	err := srv.router.Run(portStr)
	if err != nil {
		return err
	}
	return nil
}
