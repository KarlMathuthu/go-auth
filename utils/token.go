package utils

import (
    "time"

    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key") // Replace with a strong secret

type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}

func GenerateToken(email string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour) // 1 day expiry
    claims := &Claims{
        Email: email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}