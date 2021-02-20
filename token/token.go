package token

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	tokenDuration = 15
	expireOffset  = 3600
	secret        = "i'm a handsome boy"
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
func ParseToken(tokenString string) bool {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		// check signing method
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println(token)
		return false
	}

	_, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		return true
	}

	fmt.Println(token)
	return false
}
