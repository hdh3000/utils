package sendjobmail

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

// Ripped off from this: https://developers.google.com/gmail/api/quickstart/go
// worthwhile to factor out into a more general library (that saves oauth things in a dir)

// Retrieve a token, saves the token, then returns the generated client.
func GetClient() *http.Client {
	home := os.Getenv("HOME")
	tokFile := path.Join(home, ".config/sendjobmail/token.json")

	oauthConfigFile := path.Join(home, ".config/sendjobmail/client_id.json")

	f, err := ioutil.ReadFile(oauthConfigFile)
	if err != nil {
		log.Fatalf("failed to open oauth config...%s", err)
	}

	config, err := google.ConfigFromJSON(f, gmail.GmailInsertScope, gmail.GmailSendScope, gmail.GmailComposeScope)
	if err != nil {
		log.Fatalf("failed to open oauth config...%s", err)
	}

	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}

	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	json.NewEncoder(f).Encode(token)
}
