package core

import "context"

type (
	// Auth auth information table(using email or username to login)
	Auth struct {
		ID       int    `json:"-" gorm:"primary_key"`
		UserID   int64  `json:"userid"`
		Username string `json:"username"`
		Password string `json:"-"`
		Email    string `json:"email"`
	}

	// AuthStore auth to app api operations
	AuthStore interface {
		// Login login app using auth way(username/password or email/password)
		Login(ctx context.Context, username string, password string) error
		// Find find a specified user by id
		Find(context.Context, int64) (*Auth, error)
		// Count count user total number
		Count(context.Context) (int64, error)
		// Delete delete a user from auth table
		Delete(context.Context, int64) error
	}
)
