package file

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path"
)

const (
	limitSize        = 32 << 20  // limit single file size 32M
	readBuff         = 512       // read buffer size
	uploadDir string = "uploads" // upload directory
)

var (
	errOverLimitSize = errors.New("file size is too large")
	errNoSupportType = errors.New("file type no support")
)

// SaveDisk save file to disk
func SaveDisk(file *multipart.FileHeader, dir string) error {
	if !isLimit(file.Size) {
		return errOverLimitSize
	}

	if dir != "" {
		// create directory if custom directory
		if err := os.MkdirAll(path.Join(uploadDir, dir), os.ModePerm); err != nil {
			return err
		}
	}

	// save to disk
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	// read file buffer
	buff := make([]byte, readBuff)
	_, err = f.Read(buff)
	if err != nil {
		return err
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	// file Type check
	if err := fileType(buff); err != nil {
		return err
	}

	// create fd
	fd, err := os.Create(path.Join(path.Join(uploadDir, dir), file.Filename))

	if err != nil {
		return err
	}

	defer fd.Close()

	_, err = io.Copy(fd, f)
	if err != nil {
		return err
	}

	return nil
}

// isLimit check file size
func isLimit(size int64) bool {
	return size < limitSize
}

// baseDir get base Directory
func baseDir() string {
	dir, err := os.Getwd()

	if err != nil {
		return ""
	}

	return dir
}

func fileType(buff []byte) error {
	// http Detect content Type
	ft := http.DetectContentType(buff)
	switch ft {
	case "image/jpeg",
		"image/jpg",
		"image/gif",
		"image/png",
		"application/pdf",
		"text/plain; charset=utf-8",
		"application/x-gzip",
		"application/zip",
		"application/octet-stream":
		return nil
	default:
		return errNoSupportType
	}
}
