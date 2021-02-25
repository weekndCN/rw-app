package docker

import (
	"testing"
)

func TestList(t *testing.T) {
	_, err := List()
	if err != nil {
		t.Error(err)
	}
}
