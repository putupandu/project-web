package utils


import (
"time"


"github.com/golang-jwt/jwt/v5"
)

//
var SECRET = []byte("SECRET_KEY_E_LIBRARY")


func GenerateToken(userID int, duration time.Duration) (string, error) {
claims := jwt.MapClaims{
"user_id": userID,
"exp": time.Now().Add(duration).Unix(),
}


token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
return token.SignedString(SECRET)
}