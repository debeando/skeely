package flags

import (
	"os"

	"skeely/common"
)

type Flags struct {
	Files             string
	GitHubComment     bool
	GitHubPullRequest int
	GitHubRepository  string
	GitHubToken       string
	Ignore            string
	Path              string
}

var instance *Flags

func GetInstance() *Flags {
	if instance == nil {
		instance = &Flags{}
	}
	return instance
}

func (f *Flags) Load() {
	f.Files = os.Getenv("INPUT_FILES")
	f.GitHubComment = common.StringToBool(os.Getenv("INPUT_COMMENT"))
	f.GitHubPullRequest = common.StringToInt(os.Getenv("INPUT_PULLREQUEST"))
	f.GitHubRepository = os.Getenv("INPUT_REPOSITORY")
	f.GitHubToken = os.Getenv("INPUT_TOKEN")
	f.Ignore = os.Getenv("INPUT_IGNORE")
	f.Path = os.Getenv("INPUT_PATH")
}
