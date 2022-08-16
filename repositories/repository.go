package repositories

import "github.com/guillermo-st/trip-agency-api/models"

type DriverRepository interface {
	GetDrivers(pageSize uint, pageNum uint) ([]models.User, error)
	GetDriversByStatus(pageSize uint, pageNum uint, isOnTrip bool) ([]models.User, error)
	AddDriver(driver models.User) error
}

type TripRepository interface {
	StartNewTripForDriver(driverId uint) error
	FinishTripForDriver(driverId uint) error
}

type UserRepository interface {
	GetUserByEmail(email string) (models.User, error)
}
