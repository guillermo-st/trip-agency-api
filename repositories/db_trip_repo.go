package repositories

import (
	"errors"
	"time"

	"github.com/guillermo-st/trip-agency-api/models"
	"github.com/jmoiron/sqlx"
)

type DBTripRepository struct {
	db *sqlx.DB
}

func NewDBTripRepository(db *sqlx.DB) (*DBTripRepository, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	repo := &DBTripRepository{db}
	return repo, nil
}

func (repo *DBTripRepository) StartNewTripForDriver(driver models.User) error {
	var lastTrip models.Trip
	alreadyOnTripErr := repo.db.Get(&lastTrip, "SELECT trip_id, user_id, started_at, finished_at FROM trips WHERE user_id = ? AND finished_at IS NULL", driver.Id)
	if alreadyOnTripErr != nil {
		return errors.New("Can't start a new trip for a driver whose current trip has not finished.")
	}

	var err error = nil
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Panicked while starting DB transaction")
		}
	}()

	tx := repo.db.MustBegin()
	tx.MustExec("INSERT INTO trips (user_id, started_at, finished_at) VALUES($1, $2, $3);", driver.Id, time.Now(), nil)
	err = tx.Commit()
	return err
}

func (repo *DBTripRepository) FinishTripForDriver(driver models.User) error {
	var lastTrip models.Trip
	alreadyOnTripErr := repo.db.Get(&lastTrip, "SELECT trip_id, user_id, started_at, finished_at FROM trips WHERE user_id = ? AND finished_at IS NULL", driver.Id)
	if alreadyOnTripErr == nil {
		return errors.New("Can't finish a trip for a driver that is not currently on trip.")
	}

	var err error = nil
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Panicked while starting DB transaction")
		}
	}()

	tx := repo.db.MustBegin()
	tx.MustExec("UPDATE trips SET finished_at = ? WHERE user_id = ? AND finished_at IS NULL;", time.Now(), driver.Id)
	err = tx.Commit()
	return err
}
