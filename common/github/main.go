package github

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"skeely/common"
	"skeely/flags"
	"skeely/linter"
)

type GitHub struct {
	Path        string
	Enable      bool
	Token       string
	Repository  string
	PullRequest int
	Comment     string
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

func (gh *GitHub) Clear() {
	gh.Path = ""
	gh.Enable = false
	gh.Token = ""
	gh.Repository = ""
	gh.PullRequest = 0
}

func (gh *GitHub) GetFlags() {
	f := flags.GetInstance()
	gh.Clear()
	gh.Path = f.Path
	gh.Enable = f.GitHubComment
	gh.Token = f.GitHubToken
	gh.Repository = f.GitHubRepository
	gh.PullRequest = f.GitHubPullRequest
}

func (gh *GitHub) GetInputs() {
	gh.Clear()
	gh.Path = os.Getenv("INPUT_PATH")
	gh.Enable = common.StringToBool(os.Getenv("INPUT_COMMENT"))
	gh.Token = os.Getenv("INPUT_TOKEN")
	gh.Repository = os.Getenv("INPUT_REPOSITORY")
	gh.PullRequest = common.StringToInt(os.Getenv("INPUT_PULLREQUEST"))
}

func (gh *GitHub) IsSet() bool {
	if common.StringIsEmpty(gh.Path) {
		return false
	}
	if common.StringIsEmpty(gh.Token) {
		return false
	}
	if common.StringIsEmpty(gh.Repository) {
		return false
	}
	if gh.PullRequest == 0 {
		return false
	}
	if ! gh.Enable {
		return false
	}

	return true
}

func (gh *GitHub) OnTerminal() bool {
	gh.GetFlags()
	return gh.IsSet()
}

func (gh *GitHub) OnActions() bool {
	gh.GetInputs()
	return gh.IsSet()
}

func (gh *GitHub) BuildMessage() {
	l := linter.GetInstance()
	gh.Comment = "# Skeely summary\\n"
	for _, r := range l.Summary {
		gh.Comment += fmt.Sprintf("Result of file: `%s`\\n", r.File)
		for _, m := range r.Messages {
			gh.Comment += fmt.Sprintf("- **[%d]** %s\\n", m.Code, m.Message)
			
		}
	}

	fmt.Println(gh.Comment)
}

func (gh *GitHub) PushComment() error {
	requestURL := fmt.Sprintf("https://api.github.com/repos/%s/issues/%d/comments", gh.Repository, gh.PullRequest)
	jsonBody := []byte(fmt.Sprintf(`{"body": "%s"}`, gh.Comment))
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", gh.Token))

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
