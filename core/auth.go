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
		// Login to system
		Login(ctx context.Context, username string, password string) error
		// Find find a specified user by id
		Find(context.Context, int64) (*Auth, error)
		// Count count user total number
		Count(context.Context) (int64, error)
		// Delete delete a user from auth table
		Delete(context.Context, int64) error
		// Create create a new user
		Create(context.Context, *Auth) error
	}
	// AuthService auth service
	AuthService interface {
		// Login to system
		Login() error
		// logout from system
		Logout() error
	}
)
