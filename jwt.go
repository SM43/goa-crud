package blogapi

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/mitchellh/mapstructure"
)

// Claims Object to decode JWT
type Claims struct {
	Authorized bool   `json:"authorized"`
	ID         int    `json:"id"`
	Username   string `json:"username"`
}

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

func VerifyJWT(token string) error {

	jwtSecretKey, exists := os.LookupEnv("JWT_SECRET_KEY")
	if !exists {
		log.Println("Environment variable not found")
	}
	var c Claims
	parsedToken, _ := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Failed to Decode JWT")
		}
		return []byte(jwtSecretKey), nil
	})
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		mapstructure.Decode(claims, &c)
	} else {
		return errors.New("Invalid JWT")
	}
	return nil
}
