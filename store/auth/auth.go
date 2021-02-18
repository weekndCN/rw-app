package auth

import (
	"context"

	"github.com/weekndCN/rw-app/core"
	"github.com/weekndCN/rw-app/store/db"
)

// New return a auth data store instance
func New(db *db.DB) core.AuthStore {
	return &authStore{db}
}

type authStore struct {
	db *db.DB
}

// Login auth to app system
func (auth *authStore) Login(ctx context.Context, username string, password string) error {
	//out := &core.Auth{Username: username, Password: password}
	//auth.db.AutoMigrate(&core.Auth{})
	return nil
}

// Find find a suer
func (auth *authStore) Find(ctx context.Context, id int64) (*core.Auth, error) {
	return nil, nil
}

// Count count the number of auth table's user
func (auth *authStore) Count(ctx context.Context) (int64, error) {
	return 0, nil
}

// Delete delete a user from auth table
func (auth *authStore) Delete(ctx context.Context, id int64) error {
	return nil
}
