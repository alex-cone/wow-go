package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
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

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func runDocs() []string {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := "1_R_0aujAdrKGl8NGwdG-3EORBRLwJuRmP5sVEy_whTU"
	readRange := "A2:C16"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	var count string
	var parsed []string
	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		fmt.Println("Found Data")
		for _, row := range resp.Values {
			names := fmt.Sprintf("%v", row)
			names = strings.Trim(names, "[")
			names = strings.Trim(names, "]")
			names = strings.Trim(names, "/*")
			names = strings.Replace(names, "/", "", -1)
			names = strings.Trim(names, "6")
			names = strings.Trim(names, "7")
			names = strings.Trim(names, "8")
			names = strings.Trim(names, "9")
			names = strings.Trim(names, "L")
			names = strings.Replace(names, "Group", "Left", 1)
			names = strings.TrimSpace(names)
			if strings.Index(names, "R") == 0 {
				names = strings.Replace(names, "R", "Right", 1)
			}
			if strings.Contains(names, "Left") {
				temp := strings.Split(names, " ")
				count = strings.TrimSpace(temp[1])
			} else if strings.Contains(names, "Right") {
				names = names[0:5] + " " + count + "\n" + strings.Replace(names[6:], " ", "\n", 1)
			} else {
				names = strings.Replace(names, " ", "\n", 1)
			}

			if strings.Contains(names, "(") {
				names = strings.Replace(names, names[strings.Index(names, "("):strings.Index(names, ")")+1], "", 1)
			}
			finalSplit := strings.Split(names, "\n")
			parsed = append(parsed, finalSplit...)
		}
	}
	return parsed
}
