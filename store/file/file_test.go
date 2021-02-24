package file

import (
	"context"
	"testing"

	"github.com/weekndCN/rw-app/core"
	"github.com/weekndCN/rw-app/store/dbtest"
)

func TestFile(t *testing.T) {
	db, err := dbtest.Open()

	if err != nil {
		t.Error(err)
	}

	// create empty table
	err = db.Conn.AutoMigrate(&core.File{})
	if err != nil {
		t.Error(err)
	}
}

/*
func TestSaveDisk(t *testing.T) {
	file := &multipart.FileHeader{
		Filename: "test",
		Size:     20,
	}

	err := SaveDisk(file, "test")

	if err != nil {
		t.Error(err)
	}
}
*/

func TestFileList(t *testing.T) {
	db, err := dbtest.Open()

	if err != nil {
		t.Error(err)
	}

	var files []core.File
	res := db.Conn.Table("files").Find(&files)

	if res.Error != nil {
		t.Error(res.Error)
	}
}

func TestFilter(t *testing.T) {
	var ctx context.Context
	filter := &core.File{
		Name: "weeknd",
	}

	db, err := dbtest.Open()

	if err != nil {
		t.Error(err)
	}

	fs := New(db)

	_, err = fs.Find(ctx, filter)

	if err != nil {
		t.Error(err)
	}

}
