package blogapi

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT will return JWT string
func GenerateJWT(id int, username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nbf":      time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"id":       int(id),
		"username": string(username),
		"iat":      time.Now().Unix(),
		"scopes":   []string{"api:read", "api:write"},
	})

	Key, exists := os.LookupEnv("JWT_SECRET_KEY")

	if !exists {
		log.Println("Environment variable not found")
	}

	jwtToken, err := token.SignedString([]byte(Key))
	if err != nil {
		log.Println(err)
	}
	return jwtToken
}

// VerifyJWT will verifies user's JWT
func VerifyJWT(token string) error {
	jwtSecretKey, exists := os.LookupEnv("JWT_SECRET_KEY")
	if !exists {
		return errors.New("JWT Secret key not found")
	}
	splitToken := strings.Split(token, "Bearer ")
	reqToken := splitToken[1]

	parsedToken, _ := jwt.Parse(reqToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Failed to Decode JWT")
		}
		return []byte(jwtSecretKey), nil
	})
	if _, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		fmt.Println("Valid JWT")
		return nil
	}
	return errors.New("Invalid JWT")
}
