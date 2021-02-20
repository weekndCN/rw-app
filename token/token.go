package token

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	tokenDuration = 60
	expireOffset  = 3600
	secret        = "i'm a handsome boy"
)

var (
	errParseToken   = errors.New("token string parse failed")
	errInvalidToken = errors.New("token string is invalid")
)

// Claims custome claims
type Claims struct {
	Sub int64
	jwt.StandardClaims
}

// CreateToken .
func CreateToken(uid int64, username string) (string, error) {
	customClaims := &Claims{
		Sub: uid,
		StandardClaims: jwt.StandardClaims{
			Issuer:    username,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * tokenDuration).Unix(),
		},
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	// signed with token
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken parse token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		// check signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// extractToken extract token from http request header
func extractToken(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	if bearer == "" {
		bearer = r.FormValue("access_token")
	}
	return strings.TrimPrefix(bearer, "Bearer ")
}

// Middleware jwt middleware
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := extractToken(r)
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("request is Unauthorized"))
			return
		}

		_, err := ParseToken(token)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		next.ServeHTTP(w, r)
	})
}
