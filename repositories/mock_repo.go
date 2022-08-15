package repositories

import (
	"errors"

	"github.com/guillermo-st/trip-agency-api/models"
)

type MockRepository struct {
	users            map[string]models.User
	areDriversOnTrip map[string]bool
}

func NewMockRepository(initialUsers ...models.User) MockRepository {
	users := make(map[string]models.User)
	areDriversOnTrip := make(map[string]bool)
	for _, user := range initialUsers {
		users[user.Email] = user
		areDriversOnTrip[user.Email] = false
	}

	repo := MockRepository{users, areDriversOnTrip}
	return repo
}

func (repo *MockRepository) GetUserByEmail(email string) (*models.User, error) {
	user, present := repo.users[email]
	if !present {
		return nil, errors.New("User not found.")
	}
	return &user, nil
}

func (repo *MockRepository) StartNewTripForDriver(driver models.User) error {
	_, present := repo.areDriversOnTrip[driver.Email]
	if !present {
		return errors.New("Can't start trip for unexisting driver")
	}

	repo.areDriversOnTrip[driver.Email] = true
	return nil
}

func (repo *MockRepository) FinishTripForDriver(driver models.User) error {
	_, present := repo.areDriversOnTrip[driver.Email]
	if !present {
		return errors.New("Can't finish trip for unexisting driver")
	}

	repo.areDriversOnTrip[driver.Email] = false
	return nil
}

func (repo *MockRepository) AddDriver(driver models.User) error {
	_, present := repo.users[driver.Email]
	if present {
		return errors.New("Driver already exists!")
	}

	repo.users[driver.Email] = driver
	repo.areDriversOnTrip[driver.Email] = false
	return nil
}

func (repo *MockRepository) GetDrivers(pageSize uint, pageNum uint) ([]models.User, error) {
	users := make([]models.User, 0, len(repo.users))
	for _, user := range repo.users {
		users = append(users, user)
	}
	return users, nil
}

func (repo *MockRepository) GetDriversByStatus(pageSize uint, pageNum uint, isOnTrip bool) ([]models.User, error) {
	users := make([]models.User, 0, len(repo.users))
	for userEmail, userOnTrip := range repo.areDriversOnTrip {
		if userOnTrip == isOnTrip {
			users = append(users, repo.users[userEmail])
		}
	}
	return users, nil
}
