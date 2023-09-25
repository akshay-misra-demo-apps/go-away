package utils

import (
	"fmt"
	"log"
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

// IsValidPassword return true if password matched, else false
func IsValidPassword(original, hashed string) bool {
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
	claims["sub"] = user.Email
	// Token expires in 24 hours
	claims["exp"] = time.Now().Add(time.Second * 30).Unix()

	// Sign the token with a secret
	secret := []byte(SECRET)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateTokenWithRefresh(email string,
	firstName string,
	userType string,
	uid string) (signedToken string, signedRefreshToken string, err error) {

	claims := &models.SignedDetails{
		Email:     email,
		FirstName: firstName,
		Uid:       uid,
		UserType:  userType,
		MapClaims: jwt.MapClaims{
			"exp": time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
			"iat": time.Now().Local().Unix(),
			"iss": "go-cart-auth",
			"sub": email,
		},
	}

	refreshClaims := &models.SignedDetails{
		MapClaims: jwt.MapClaims{
			"exp": time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET))

	if err != nil {
		log.Panic(err)
		return
	}

	return token, refreshToken, err
}

func ValidateToken(tokenString string) (bool, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return the secret key used for signing
		return []byte("go-cart.com.007"), nil
	})

	if err != nil {
		return false, err
	}

	// Check if the token is valid
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	} else {
		return false, fmt.Errorf("invalid token")
	}
}
