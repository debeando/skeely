package github

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Custom HTTP client for this module.
var Client HTTPClient

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func init() {
	Client = &http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
}

func Comment(token, repository string, pullRequest int, comment string) error {
	requestURL := fmt.Sprintf("https://api.github.com/repos/%s/issues/%d/comments", repository, pullRequest)
	jsonBody := []byte(fmt.Sprintf(`{"body": "%s"}`, comment))
 	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := Client.Do(req)
	if err != nil {
		return err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.New(fmt.Sprintf("client: could not read response body: %s\n", err))
	}

	if res.StatusCode != http.StatusCreated {
		return errors.New(string(resBody))
	}

	return nil
}
