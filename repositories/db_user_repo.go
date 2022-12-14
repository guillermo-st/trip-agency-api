package repositories

import (
	"github.com/guillermo-st/trip-agency-api/models"
	"github.com/jmoiron/sqlx"
)

type DBUserRepository struct {
	db *sqlx.DB
}

func NewDBUserRepository(db *sqlx.DB) *DBUserRepository {
	return &DBUserRepository{db}
}

func (repo *DBUserRepository) GetUserByEmail(email string) (models.User, error) {
	var usr models.User
	err := repo.db.Get(&usr, "SELECT user_id, first_name, last_name, email, password, role_id FROM users WHERE email = $1", email)
	return usr, err
}
