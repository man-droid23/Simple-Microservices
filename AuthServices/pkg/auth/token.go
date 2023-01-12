package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

var jwtSecret = []byte(os.Getenv("SECRETKEY"))

func GerateJWTToken(claims *jwt.MapClaims) (string, error) {

}
