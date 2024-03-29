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

func CreateToken(email string, id string) (string, error) {

	fmt.Println("JWT_SECRET_KEY:", os.Getenv("JWT_SECRET_KEY"))
	// Create a new claim

	claims := jwt.MapClaims{
		"email": email,
		"id":    id,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))

	return tokenString, err
}

func ParseToken(tokenString string) (interface{}, interface{}, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil

	})

	if err != nil {
		return nil, nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil, fmt.Errorf("Invalid token")
	}

	email := claims["email"]
	userId := claims["id"]
	return email, userId, err
}

type TokenResponse struct {
	Id    string
	Email string
}
