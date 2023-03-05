package flags

import (
	"flag"

	"skeely/command/help"
)

type Flags struct {
	Help              bool
	Path              string
	Version           bool
	Git               bool
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
	fGit := flag.Bool("git", false, "")
	fGitHubComment := flag.Bool("github-comment", false, "")
	fGitHubPullRequest := flag.Int("github-pull-request", 0, "")
	fGitHubRepository := flag.String("github-repository", "", "")
	fGitHubToken := flag.String("github-token", "", "")
	flag.Usage = func() { help.Show(1) }
	flag.Parse()

	f.Help = *fHelp
	f.Path = *fPath
	f.Version = *fVersion
	f.Git = *fGit
	f.GitHubComment = *fGitHubComment
	f.GitHubPullRequest = *fGitHubPullRequest
	f.GitHubRepository = *fGitHubRepository
	f.GitHubToken = *fGitHubToken
}
