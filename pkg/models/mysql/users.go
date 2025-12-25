package mysql

import (
	"database/sql"

	"kdp.net/snippetbox/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

// We'll use the Insert method to add a new record to the users table.
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// We'll use the Authenticate method to verify the credentials passed to it.
// This method returns the unique ID of the associated user if they're found,
// or a ErrInvalidCredentials error otherwise.
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

// We'll use the Get method to fetch a specific record from the users table.
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}