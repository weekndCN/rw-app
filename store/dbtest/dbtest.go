package dbtest

import (
	"fmt"

	"github.com/weekndCN/rw-app/store/db"
)

// Open a connection
func Open() (*db.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"root",
		"123456",
		"localhost",
		"test",
	)
	return db.Open("mysql", dsn)
}
