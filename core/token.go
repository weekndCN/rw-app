package core

import "net/http"

type (
	// Token BearToekn
	Token interface {
		// create a JWT
		Create(http.ResponseWriter, *Auth) error
		// Delete a JWT
		Delete() error
		// Get a jwt from request
		Get(*http.Request) (*Auth, error)
	}
)
