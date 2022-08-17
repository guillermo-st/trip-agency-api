package repositories

import (
	"errors"
	"fmt"

	"github.com/guillermo-st/trip-agency-api/models"
	"github.com/jmoiron/sqlx"
)

type DBDriverRepository struct {
	db *sqlx.DB
}

func NewDBDriverRepository(db *sqlx.DB) *DBDriverRepository {
	return &DBDriverRepository{db}
}

func (repo *DBDriverRepository) GetDrivers(pageSize uint, pageNum uint) ([]models.User, error) {
	var drivers []models.User

	query := `SELECT user_id, first_name, last_name, email, password, role_id 
				FROM users 
				WHERE role_id = $1 
				ORDER BY user_id ASC
				LIMIT $2 OFFSET $3;`

	offset := pageNum * pageSize

	err := repo.db.Select(&drivers, query, models.DRIVER_ROLE_ID, pageSize, offset)
	if err != nil {
		return nil, err
	}
	return drivers, nil
}

func (repo *DBDriverRepository) GetDriversByStatus(pageSize uint, pageNum uint, isOnTrip bool) ([]models.User, error) {

	var drivers []models.User

	queryFinishedAt := "IS NULL"
	query := `SELECT u.user_id as user_id, u.first_name as first_name, u.last_name as last_name, u.email as email, u.password as password, u.role_id as role_id
				FROM users u INNER JOIN trips t ON t.user_id = u.user_id
				WHERE t.finished_at %v AND u.role_id = $1 AND t.trip_id >= ALL (
					SELECT tr.trip_id FROM trips tr WHERE tr.user_id = u.user_id 
				)
				ORDER BY user_id ASC
				LIMIT $2 OFFSET $3;`

	if !isOnTrip {
		noTripsQuery := `SELECT u.user_id, u.first_name, u.last_name, u.email, u.password, u.role_id 
					FROM users u
					WHERE NOT EXISTS
					(SELECT t.trip_id FROM trips t WHERE t.user_id = u.user_id)
					AND u.role_id = $1 
					UNION `

		query = noTripsQuery + query
		queryFinishedAt = "IS NOT NULL"
	}
	query = fmt.Sprintf(query, queryFinishedAt)
	offset := pageNum * pageSize
	err := repo.db.Select(&drivers, query, models.DRIVER_ROLE_ID, pageSize, offset)
	return drivers, err
}

func (repo *DBDriverRepository) AddDriver(driver models.User) error {
	var err error = nil
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Panicked while starting DB transaction")
		}
	}()

	tx := repo.db.MustBegin()
	tx.MustExec("INSERT INTO users (first_name, last_name, email, password, role_id) VALUES($1, $2, $3, $4, $5);", driver.FirstName, driver.LastName, driver.Email, driver.Password, models.DRIVER_ROLE_ID)
	err = tx.Commit()
	return err
}
