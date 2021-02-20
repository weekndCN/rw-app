package auth

import (
	"context"
	"errors"

	"github.com/weekndCN/rw-app/core"
	"github.com/weekndCN/rw-app/store/db"
)

var (
	errNotFound    = errors.New("Not found user")
	errExistUser   = errors.New("User exist in database already")
	errInvalidAuth = errors.New("auth information is wrong")
)

type authStore struct {
	db *db.DB
}

// New return a auth data store instance
func New(db *db.DB) core.AuthStore {
	return &authStore{db}
}

// Login auth to app system
func (auth *authStore) Login(ctx context.Context, username string, password string) (*core.Auth, error) {
	// SELECT * FROM auths where password=<passwor> and (username=<username> or email="username")
	var user core.Auth
	res := auth.db.Conn.Table("auths").Where("password=? and username=?", password, username).Or("password=? and email=?", password, username).Scan(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errNotFound
	}
	return &user, nil
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

// Create a new user to auth table
func (auth *authStore) Create(ctx context.Context, register *core.Auth) error {
	res := auth.db.Conn.Where("username=?", register.Username).Or("email=?", register.Email).Find(&core.Auth{})

	if res.RowsAffected == 0 {
		reg := auth.db.Conn.Create(register)
		if reg.Error != nil {
			return reg.Error
		}

		return nil

	}

	return errExistUser
}
