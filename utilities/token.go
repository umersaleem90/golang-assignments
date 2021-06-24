package utilities

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type claims struct {
	ID uint `json:"ID"`
	jwt.StandardClaims
}

var jwtKey = []byte("secret_key")

func CreateToken(userID uint) (*string, error){
	expirationTime := time.Now().Add(time.Hour * 72)
	claims := &claims{
		ID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func VerifyToken(r *http.Request) (context.Context, error){
	var tokenString string
	reqToken := strings.Split(r.Header.Get("Authorization"), " ")
	if len(reqToken) > 1 {
		tokenString = reqToken[1]
	}
	claims := &claims{}
	tkn, err := jwt.ParseWithClaims(tokenString, claims, 
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
			
	if err != nil {
		return nil, err
	}
	
	ctx := context.WithValue(r.Context(), "userID", claims.ID) // adding the user ID to the context

	if tkn.Valid {
		return ctx, nil
	} else {
		return nil, errors.New("invalid token")
	}
}