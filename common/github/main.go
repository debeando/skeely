package github

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"skeely/flags"
	"skeely/message"
)

type GitHub struct {
	CommentText string
}

var Client HTTPClient

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func init() {
	Client = &http.Client{
		Timeout: time.Duration(10) * time.Second,
	}
}

func (gh *GitHub) IsSet() bool {
	return false
}

func (gh *GitHub) Comment(messages message.Plugins) {
	if !messages.IsSet() {
		return
	}

	gh.CommentText = "# Skeely summary:\\n"
	gh.CommentText += "Is a Schema Linter for MySQL, this tool help to identifying some common and uncommon mistakes on data model.\\n\\n"
	for _, r := range messages {
		gh.CommentText += fmt.Sprintf("**Result of file:** `%s`\\n", r.File)
		for _, m := range r.Messages {
			gh.CommentText += fmt.Sprintf("- **[%d]** %s\\n", m.Code, m.Message)
		}
		if len(r.Messages) == 0 {
			gh.CommentText += "- Looks ok.\\n\\n"
		} else {
			gh.CommentText += "\\n"
		}
	}

	gh.CommentText += "For more help, plese visit: https://github.com/debeando/skeely"
}

func (gh *GitHub) Push() error {
	f := flags.GetInstance()

	if !f.GitHubComment {
		return nil
	}

	if len(gh.CommentText) == 0 {
		return nil
	}

	requestURL := fmt.Sprintf("https://api.github.com/repos/%s/issues/%d/comments", f.GitHubRepository, f.GitHubPullRequest)
	jsonBody := []byte(fmt.Sprintf(`{"body": "%s"}`, gh.CommentText))
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", f.GitHubToken))

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
