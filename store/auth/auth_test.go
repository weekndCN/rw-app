package auth

import (
	"context"
	"testing"

	"github.com/weekndCN/rw-app/core"
	"github.com/weekndCN/rw-app/store/dbtest"
)

var noContext = context.TODO()

func TestAuth(t *testing.T) {
	db, err := dbtest.Open()

	if err != nil {
		t.Error(err)
	}

	// create empty table
	err = db.Conn.AutoMigrate(&core.Auth{})
	if err != nil {
		t.Error(err)
	}
	//_ := New(conn).(*authStore)
}
