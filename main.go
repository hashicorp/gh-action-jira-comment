package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/tomhjp/gh-action-jira/config"
	"github.com/tomhjp/gh-action-jira/jira"
)

func main() {
	err := comment()
	if err != nil {
		log.Fatal(err)
	}
}

func comment() error {
	issue := os.Getenv("INPUT_ISSUE")
	if issue == "" {
		return errors.New("no issue provided as input")
	}
	comment := os.Getenv("INPUT_COMMENT")

	config, err := config.ReadConfig()
	if err != nil {
		return err
	}

	err = addComment(config, issue, comment)
	if err != nil {
		return err
	}

	fmt.Printf("Commented on issue %s\n", issue)

	return nil
}

func addComment(config config.JiraConfig, issue, comment string) error {
	payload := struct {
		Body string `json:"body"`
	}{
		Body: comment,
	}
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// Use the REST API v2 because it has a much simpler schema for body (just a string instead of ADF).
	_, err = jira.DoRequest(config, "POST", fmt.Sprintf("/rest/api/2/issue/%s/comment", issue), url.Values{}, bytes.NewReader(reqBody))
	if err != nil {
		indentedBody, marshalErr := json.MarshalIndent(payload, "", "  ")
		if marshalErr != nil {
			// We made a best effort, oh well, just print it ugly.
			indentedBody = reqBody
		}
		fmt.Println("Request body:")
		fmt.Printf("%s\n", string(indentedBody))
		return err
	}

	return nil
}
