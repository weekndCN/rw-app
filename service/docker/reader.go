package docker

import (
	"bytes"
	"encoding/binary"
	"io"
)

// Read reader data
func Read(logReadCloser io.ReadCloser, p []byte) (n int, err error) {
	buffer := &bytes.Buffer{}
	lastHeader := make([]byte, 8)
	if buffer.Len() > 0 {
		return logReadCloser.Read(p)
	}

	buffer.Reset()
	_, err = logReadCloser.Read(lastHeader)
	if err != nil {
		return 0, err
	}
	count := binary.BigEndian.Uint32(lastHeader[4:])
	_, err = io.CopyN(buffer, logReadCloser, int64(count))
	if err != nil {
		return 0, err
	}
	return buffer.Read(p)
}
