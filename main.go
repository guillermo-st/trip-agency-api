package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guillermo-st/trip-agency-api/controllers"
	"github.com/guillermo-st/trip-agency-api/middleware"
	"github.com/guillermo-st/trip-agency-api/repositories"
	"github.com/guillermo-st/trip-agency-api/services"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=local sslmode=disable")
	//db, err := sqlx.Open("postgres", "host=<yourhost> port=<yourport> user=<youruser> password=<yourpassword> dbname=<yourdbname> sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	driverRepo, err := repositories.NewDBDriverRepository(db)
	if err != nil {
		panic(err)
	}
	driverCtrl := controllers.NewDriverController(driverRepo)

	userRepo, err := repositories.NewDBUserRepository(db)
	if err != nil {
		panic(err)
	}

	authServ, err := services.NewJsonWebTokenService()
	if err != nil {
		panic(err)
	}

	loginCtrl := controllers.NewLoginController(userRepo, *authServ)

	router := gin.Default()

	router.GET("/drivers", middleware.AuthorizeJWT(true), driverCtrl.GetDrivers)
	router.GET("/drivers-by-status", middleware.AuthorizeJWT(true), driverCtrl.GetDriversByStatus)
	router.POST("/drivers", middleware.AuthorizeJWT(true), driverCtrl.AddDriver)

	router.POST("/login", loginCtrl.Login)

	err = router.Run(":8000")
	if err != nil {
		panic(err)
	}
}
