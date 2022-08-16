package repositories

import (
	"errors"

	"github.com/guillermo-st/trip-agency-api/models"
	"github.com/jmoiron/sqlx"
)

type DBDriverRepository struct {
	db *sqlx.DB
}

func NewDBDriverRepository(db *sqlx.DB) (*DBDriverRepository, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	repo := &DBDriverRepository{db}
	return repo, nil
}

func (repo *DBDriverRepository) GetDrivers(pageSize uint, pageNum uint) ([]models.User, error) {
	var drivers []models.User

	query := `SELECT user_id, first_name, last_name, email, password, role_id 
				FROM users 
				WHERE role_id = 2
				ORDER BY user_id ASC
				LIMIT $1 OFFSET $2;`

	offset := pageNum * pageSize

	err := repo.db.Select(&drivers, query, pageSize, offset)
	if err != nil {
		return nil, err
	}
	return drivers, nil
}

func (repo *DBDriverRepository) GetDriversByStatus(pageSize uint, pageNum uint, isOnTrip bool) ([]models.User, error) {

	var drivers []models.User

	queryFinishedAt := "IS NULL"
	query := `SELECT u.user_id, u.first_name, u.last_name, u.email, u.password, u.role_id 
				FROM users u INNER JOIN trips t ON t.user_id = u.user_id
				WHERE t.finished_at $1 AND u.role_id = 2
				ORDER BY user_id ASC
				LIMIT $2 OFFSET $3;`

	if !isOnTrip {
		noTripsQuery := `SELECT u.user_id, u.first_name, u.last_name, u.email, u.password, u.role_id 
					FROM users u
					WHERE NOT EXISTS
					(SELECT t.trip_id FROM trips t WHERE t.user_id = u.user_id)
					AND u.role_id = 2
					UNION `

		query = noTripsQuery + query
		queryFinishedAt = "IS NOT NULL"
	}

	offset := pageNum * pageSize
	err := repo.db.Select(&drivers, query, queryFinishedAt, pageSize, offset)
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
	tx.MustExec("INSERT INTO users (first_name, last_name, email, password, role_id) VALUES($1, $2, $3, $4, $5);", driver.FirstName, driver.LastName, driver.Email, driver.Password, 2)
	err = tx.Commit()
	return err
}
