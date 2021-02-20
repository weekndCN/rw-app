package token

import (
	"fmt"
	"net/http/httptest"
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
	var uid int64 = 11
	tokenString, err := CreateToken(uid, "weeknd")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(tokenString)

	_, err = ParseToken(tokenString)
	if err != nil {
		t.Error(err)
	}

}

func TestExtractToken(t *testing.T) {
	r := httptest.NewRequest("GET", "/", nil)
	r.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJTdWIiOjExLCJleHAiOjE2MTM4MzU4MDYsImlhdCI6MTYxMzgzNDkwNiwiaXNzIjoid2Vla25kIn0.lEbHX1Y1iS_kBQxnjR0F3wiMCJHAzqmv_w4mmWqn1O0")
	token := extractToken(r)
	if token == "" {
		t.Error("extract token from request failed")
	}
}
