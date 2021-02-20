package token

import (
	"fmt"
	"testing"
)

func TestCreateToken(t *testing.T) {
	var uid int64 = 11
	_, err := CreateToken(uid, "weeknd")
	if err != nil {
		t.Error(err)
	}
}

func TestParseToken(t *testing.T) {
	var tokenString string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdWIiOjEsImV4cCI6MTYxMzgyMjY1MSwiaWF0IjoxNjEzODIxNzUxLCJpc3MiOiIxIn0.iLZAbglwwegr0CsJdVnSzxlWpx3y3_ZiiANYlxrt9ZI"
	fmt.Println(ParseToken(tokenString))
}
