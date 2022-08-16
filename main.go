package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guillermo-st/trip-agency-api/controllers"
	"github.com/guillermo-st/trip-agency-api/repositories"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Open("postgres", "host=<yourhost> port=<yourport> user=<youruser> password=<yourpassword> dbname=<yourdbname> sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	driverRepo, err := repositories.NewDBDriverRepository(db)
	if err != nil {
		panic(err)
	}

	driverCtrl := controllers.NewDriverController(driverRepo)

	router := gin.Default()

	router.GET("/drivers", driverCtrl.GetDrivers)
	router.GET("/drivers-by-status", driverCtrl.GetDriversByStatus)
	router.POST("/drivers", driverCtrl.AddDriver)

	err = router.Run(":8000")
	if err != nil {
		panic(err)
	}
}
