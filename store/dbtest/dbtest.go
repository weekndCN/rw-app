package dbtest

import (
	"fmt"
	"os"

	"github.com/weekndCN/rw-app/store/db"
)

// Open a connection
func Open() (*db.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("RWPLUS_MYSQL_USER"),
		os.Getenv("RWPLUS_MYSQL_PASSWORD"),
		os.Getenv("RWPLUS_MYSQL_ADDR"),
		os.Getenv("RWPLUS_MYSQL_DB"),
	)
	return db.Open("mysql", dsn)
}
