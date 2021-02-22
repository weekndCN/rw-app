package core

import (
	"context"
	"time"
)

type (
	// File upload file struct
	File struct {
		// file unique id
		ID int64 `json:"id"`
		// file name
		Name string `json:"name"`
		// file uploader
		User string `json:"user"`
		// file size
		Size int64 `json:"size"`
		// file location
		Location string `json:"location"`
		// file type
		Type string `json:"type"`
		// create date
		CreateAt time.Time `json:"create_at"`
	}

	// FileStore abstract file methods
	FileStore interface {
		// Create store file info to data store
		Create(ctx context.Context, file *File) error
		// Find return a file from data store
		Find(ctx context.Context) error
		// Delete delete a file from data store
		Delete(ctx context.Context) error
		// List list all files from data store
		List(ctx context.Context) (*[]File, error)
	}
)
