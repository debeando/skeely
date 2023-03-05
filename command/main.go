package command

import (
	"fmt"
	"os"

	"skeely/command/help"
	"skeely/command/version"
	"skeely/common/github"
	"skeely/config"
	"skeely/flags"
	"skeely/linter"

	_ "skeely/plugins"
)

func Run() {
	gh := github.GitHub{}
	c := config.GetInstance()
	l := linter.GetInstance()
	f := flags.GetInstance()
	f.Load()

	if gh.OnActions() {
		l.Path = gh.Path
		l.Git = gh.Git
		l.Run()
		gh.BuildMessage()
		gh.PushComment()
		os.Exit(0)
	}

	switch {
	case f.Version:
		fmt.Println(version.VERSION)
		os.Exit(0)
	case f.Help:
		help.Show(0)
	case len(f.Path) == 0:
		help.Show(1)
	}

	err := c.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	l.Path = f.Path
	l.Git = gh.Git
	l.Run()

	if gh.OnTerminal() {
		gh.BuildMessage()
		gh.PushComment()
	}

	for _, r := range l.Summary {
		fmt.Println(fmt.Sprintf("> File: %s", r.File))
		for _, m := range r.Messages {
			fmt.Println(fmt.Sprintf("- [%d] %s", m.Code, m.Message))
		}
		// TODO: Si el fichero esta bien, ponerlo.
	}
}
