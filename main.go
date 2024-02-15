package main

import (
	"fmt"
	"os"

	"skeely/common/github"
	"skeely/common/terminal"
	"skeely/config"
	"skeely/flags"
	"skeely/linter"
	"skeely/version"

	"github.com/spf13/cobra"
)

func main() {
	var f = flags.GetInstance()
	var rootCmd = &cobra.Command{
		Use: "skeely [COMMANDS] [OPTIONS]",
		Long: `skeely is a schema linter for MySQL, this tool help to identifying
some common and uncommon mistakes on data model.

For more help, plese visit: https://github.com/debeando/skeely`,
		Example: `
  # Lint directory
  $ skeely --path=assets/examples/

  # Lint specific file
  $ skeely --files=assets/examples/case01.sql

  # Lint specific file and ignore codes
  $ skeely --files=assets/examples/case01.sql --ignore=103,104,305,406

  # Lint and push summary as comment into GitHub Pull Request.
  $ skeely --path=assets/examples/case01.sql \
           --github-comment \
           --github-token=${{github.token}} \
           --github-repository=$GITHUB_REPOSITORY \
           --github-pull-request=${{github.event.pull_request.number}}
`,
		Run: func(cmd *cobra.Command, args []string) {
			gh := github.GitHub{}
			c := config.GetInstance()

			if err := c.Load(); err != nil {
				fmt.Println(err)
				os.Exit(2)
			}

			msgPlugins := linter.Run()

			gh.Comment(msgPlugins)
			gh.Push()

			terminal.Print(msgPlugins)
		},
	}

	f.Load()

	rootCmd.Flags().BoolVarP(&f.GitHubComment, "github-comment", "", false, "Send summary as comment into GitHub.")
	rootCmd.Flags().IntVar(&f.GitHubPullRequest, "github-pull-request", 0, "Pull request number.")
	rootCmd.Flags().StringVar(&f.Files, "files", "", "List of files to lint, separated by space.")
	rootCmd.Flags().StringVar(&f.GitHubRepository, "github-repository", "", "Repository path on github.")
	rootCmd.Flags().StringVar(&f.GitHubToken, "github-token", "", "Token to auth in github.")
	rootCmd.Flags().StringVar(&f.Ignore, "ignore", "", "List of error codes separated by comma to ignore.")
	rootCmd.Flags().StringVar(&f.Path, "path", "", "Path of the directory to start to find *.sql to lint.")
	rootCmd.AddCommand(version.NewCommand())
	rootCmd.Execute()
}
