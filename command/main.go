package command

import (
	"fmt"
	"os"

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
		l.Run()
		gh.BuildMessage()
		gh.PushComment()
		os.Exit(0)
	}

	switch {
	case f.Version:
		fmt.Println(VERSION)
		os.Exit(0)
	case f.Help:
		help(0)
	case len(f.Path) == 0:
		help(1)
	}

	err := c.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	l.Path = f.Path
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
	}
}
