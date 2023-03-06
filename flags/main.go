package flags

import (
	"flag"
	"os"

	"skeely/command/help"
	"skeely/common"
)

type Flags struct {
	Files             string
	GitHubComment     bool
	GitHubPullRequest int
	GitHubRepository  string
	GitHubToken       string
	Help              bool
	Path              string
	Version           bool
}

var instance *Flags

func GetInstance() *Flags {
	if instance == nil {
		instance = &Flags{}
	}
	return instance
}

func (f *Flags) Load() {
	fFiles := flag.String("files", "", "")
	fGitHubComment := flag.Bool("github-comment", false, "")
	fGitHubPullRequest := flag.Int("github-pull-request", 0, "")
	fGitHubRepository := flag.String("github-repository", "", "")
	fGitHubToken := flag.String("github-token", "", "")
	fHelp := flag.Bool("help", false, "")
	fPath := flag.String("path", "", "")
	fVersion := flag.Bool("version", false, "")
	flag.Usage = func() { help.Show(1) }
	flag.Parse()

	f.Files = *fFiles
	f.GitHubComment = *fGitHubComment
	f.GitHubPullRequest = *fGitHubPullRequest
	f.GitHubRepository = *fGitHubRepository
	f.GitHubToken = *fGitHubToken
	f.Help = *fHelp
	f.Path = *fPath
	f.Version = *fVersion

	if len(f.Files) == 0 && len(os.Getenv("INPUT_FILES")) > 0 {
		f.Files = os.Getenv("INPUT_FILES")
	}

	if len(f.Path) == 0 && len(os.Getenv("INPUT_PATH")) > 0 {
		f.Path = os.Getenv("")
	}

	if f.GitHubComment == false && common.StringToBool(os.Getenv("INPUT_COMMENT")) == true {
		f.GitHubComment = true
	}

	if f.GitHubPullRequest == 0 && common.StringToInt(os.Getenv("INPUT_PULLREQUEST")) > 0 {
		f.GitHubPullRequest = common.StringToInt(os.Getenv("INPUT_PULLREQUEST"))
	}

	if len(f.GitHubRepository) == 0 && len(os.Getenv("INPUT_REPOSITORY")) > 0 {
		f.GitHubRepository = os.Getenv("INPUT_REPOSITORY")
	}

	if len(f.GitHubToken) == 0 && len(os.Getenv("INPUT_TOKEN")) > 0 {
		f.GitHubToken = os.Getenv("INPUT_TOKEN")
	}
}
