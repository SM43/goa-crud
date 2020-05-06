package blogapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	oauth "github.com/sm43/goa-crud/gen/oauth"
)

// oauth service example implementation.
// The example methods log the requests and return zero values.
type oauthsrvc struct {
	db     *gorm.DB
	logger *log.Logger
}

// NewOauth returns the oauth service implementation.
func NewOauth(db *gorm.DB, logger *log.Logger) oauth.Service {
	return &oauthsrvc{db, logger}
}

type OAuthAccessResponse struct {
	AccessToken string `json:"access_token"`
}

type userDetails struct {
	id    int
	email string
}

var users = make([]userDetails, 0)

// Github authentication to post a new blog
func (s *oauthsrvc) Oauth(ctx context.Context, p *oauth.OauthPayload) (res string, err error) {

	reqURL := ghOAuthURLForCode(*p.Token)

	req, err := http.NewRequest(http.MethodPost, reqURL, nil)
	if err != nil {
		fmt.Fprintf(os.Stdout, "could not create HTTP request: %v", err)
	}
	req.Header.Set("accept", "application/json")

	// // Send out the HTTP request
	httpClient := http.Client{}
	result, err := httpClient.Do(req)
	if err != nil {
		println(os.Stdout, "could not send HTTP request: %v", err)
	}

	// Parse the request body into the `OAuthAccessResponse` struct
	var t OAuthAccessResponse
	if err := json.NewDecoder(result.Body).Decode(&t); err != nil {
		fmt.Fprintf(os.Stdout, "could not parse JSON response: %v", err)
	}
	s.logger.Println("User access token", t.AccessToken)

	username, id := getUserDetails(t.AccessToken)

	s.logger.Println("Username ", username)
	s.logger.Println(id)

	data := userDetails{id: id, email: username}

	users = append(users, data)
	s.logger.Println("Users", users)

	jwtToken := GenerateJWT(id, username)

	s.logger.Println("Jwt token: ", jwtToken)

	return jwtToken, nil
}

func ghOAuthURLForCode(code string) string {
	// Get the client_id and client_secret environment variable
	clientID, err := os.LookupEnv("CLIENT_ID")
	if !err {
		log.Println("Environment variable not found")
	}
	clientSecret, err := os.LookupEnv("CLIENT_SECRET")
	if !err {
		log.Println("Environment variable not found")
	}
	return fmt.Sprintf(
		"https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s",
		clientID, clientSecret, code)
}

func getUserDetails(accessToken string) (string, int) {
	httpClient := http.Client{}
	reqURL := fmt.Sprintf("https://api.github.com/user")
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	req.Header.Set("Authorization", "token "+accessToken)
	if err != nil {
		log.Print(err)
	}
	req.Header.Set("Access-Control-Allow-Origin", "*")
	req.Header.Set("accept", "application/json")

	// Send out the HTTP request
	res, err := httpClient.Do(req)
	if err != nil {
		log.Print(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var userData map[string]interface{}
	if err := json.Unmarshal(body, &userData); err != nil {
		log.Print(err)
	}

	username := userData["login"].(string)
	id := userData["id"].(float64)
	return string(username), int(id)
}
