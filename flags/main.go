package flags

import (
	"flag"
)

type Flags struct {
	Help              bool
	Path              string
	Version           bool
	GitHubComment     bool
	GitHubPullRequest int
	GitHubRepository  string
	GitHubToken       string
}

var instance *Flags

func GetInstance() *Flags {
	if instance == nil {
		instance = &Flags{}
	}
	return instance
}

func (f *Flags) Load() {
	fHelp := flag.Bool("help", false, "")
	fPath := flag.String("path", "", "")
	fVersion := flag.Bool("version", false, "")
	fGitHubComment := flag.Bool("github-comment", false, "")
	fGitHubPullRequest := flag.Int("github-pull-request", 0, "")
	fGitHubRepository := flag.String("github-repository", "", "")
	fGitHubToken := flag.String("github-token", "", "")
	flag.Parse()

	f.Help = *fHelp
	f.Path = *fPath
	f.Version = *fVersion
	f.GitHubComment = *fGitHubComment
	f.GitHubPullRequest = *fGitHubPullRequest
	f.GitHubRepository = *fGitHubRepository
	f.GitHubToken = *fGitHubToken
}
