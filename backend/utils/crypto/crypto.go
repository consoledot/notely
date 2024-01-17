package cryptolib

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var _ = godotenv.Load()

func Hash(text string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(text), 14)
	return string(bytes), err
}

func CompareHashWithText(hash string, text string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))

	return err == nil
}

func CreateToken(id string) (string, error) {
	// err := godotenv.Load()
	// if err != nil {
	// 	fmt.Println("Error loading .env file: \n", err)
	// }
	fmt.Println("JWT_SECRET_KEY:", os.Getenv("JWT_SECRET_KEY"))
	// Create a new claim
	fmt.Printf("id %v \n", id)
	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	fmt.Printf("claim %v \n", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	fmt.Printf("token %v \n", token)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	fmt.Printf("tokenstr %v \n", tokenString)

	return tokenString, err
}
