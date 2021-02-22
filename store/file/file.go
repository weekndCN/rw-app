package file

import (
	"context"

	"github.com/weekndCN/rw-app/core"
	"github.com/weekndCN/rw-app/store/db"
)

type fileStore struct {
	db *db.DB
}

// New return a file data
func New(db *db.DB) core.FileStore {
	return &fileStore{db}
}

// Upload method
func (f *fileStore) Create(ctx context.Context, file *core.File) error {
	// data store
	res := f.db.Conn.Create(&file).Table("files")
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// Find return a file
func (f *fileStore) Find(context.Context) error {
	return nil
}

//  Delete delete a file from data store
func (f *fileStore) Delete(context.Context) error {
	return nil
}

// List List files from data store
func (f *fileStore) List(context.Context) (*[]core.File, error) {
	var files []core.File
	res := f.db.Conn.Table("files").Find(&files)
	if res.Error != nil {
		return nil, res.Error
	}
	return &files, nil
}
