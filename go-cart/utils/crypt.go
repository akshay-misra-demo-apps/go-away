package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gocart.com/go-cart/models"
	"golang.org/x/crypto/bcrypt"
)

var SECRET = "go-cart.com.007"

func Encrypt(key string, cost int) (hashed string, err error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(key), cost)
	return string(hashedBytes), err
}

// Compare return true if password matched, else false
func Compare(original, hashed string) bool {
	fmt.Printf("original password: %v - %v", original, hashed)
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(original))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false
	}

	return true
}

func GenerateToken(user *models.User) (string, error) {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	// Token expires in 24 hours
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Sign the token with a secret
	secret := []byte(SECRET)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
