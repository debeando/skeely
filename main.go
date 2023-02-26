package main

import (
	"flag"
	"fmt"
	"os"

	"mylinter/common"
	"mylinter/common/github"
	"mylinter/config"
	"mylinter/directory"
	"mylinter/registry"
	"mylinter/table"

	_ "mylinter/plugins"
)

const VERSION string = "0.0.0-beta.2"
const USAGE = `mylinter %s Is a MySQL Migration Lint and this tool help to identifying some
common and uncommon mistakes data model.

Usage: mylinter [--help | --path | --version]

Options:
  --comment        Send summary as comment into GitHub.
  --help           Show this help.
  --path           Lint directory or file.
  --pull-request   Pull Request ID.
  --repository     Repository name.
  --token          Token to auth in github.
  --version        Print version numbers.

Example:

  # Lint directory
  $ mylinter --path=assets/examples/

  # Lint specific file
  $ mylinter --path=assets/examples/case01.sql

  # Lint and push summary as comment into GitHub Pull Request.
  $ mylinter --path=assets/examples/case01.sql \
             --comment \
             --token=${{github.token}} \
             --repository=$GITHUB_REPOSITORY \
             --pull-request=${{github.event.pull_request.number}}

For more help, plese visit: https://github.com/debeando/mylinter
`

var exitCode = 0

func main() {
	fHelp := flag.Bool("help", false, "")
	fVersion := flag.Bool("version", false, "")
	fComment := flag.Bool("comment", false, "")
	fPath := flag.String("path", "", "")
	fToken := flag.String("token", "", "")
	fRepository := flag.String("repository", "", "")
	fPullRequest := flag.Int("pull-request", 0, "")
	flag.Usage = func() { help(1) }
	flag.Parse()

	switch {
	case *fVersion:
		fmt.Println(VERSION)
		os.Exit(0)
	case *fHelp:
		help(0)
	case len(*fPath) == 0:
		help(1)
	}

	cnf := config.GetInstance()
	err := cnf.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	var msgComment string
	msgComment = "# mylinter summary\\n"

	directory.Explore(*fPath, func(fileName, fileContent string) {
		fmt.Println("> File:", fileName)

		tbl := table.Table{}
		if tbl.Parser(fileContent) != nil {
			return
		}

		for key := range registry.Plugins {
			if creator, ok := registry.Plugins[key]; ok {
				plugin := creator()
				messages := plugin.Run(registry.Arguments{Path: fileName, Table: tbl})

				for _, message := range messages {
					if common.IntInArrayInt(cnf.IgnoreCodes(tbl.Name), key+message.Code) {
						continue
					}

					fmt.Println(fmt.Sprintf("- [%d] %s", key+message.Code, message.Message))
					exitCode++
				}
			}
		}

		if exitCode == 0 {
			fmt.Println("  Looks ok")
		}
		fmt.Println()

		if *fComment {
			if exitCode > 0 {
				msgComment += fmt.Sprintf("Result of file: `%s`\\n", fileName)
				msgComment += "Fix follow issues:\\n"
				for key := range registry.Plugins {
					if creator, ok := registry.Plugins[key]; ok {
						plugin := creator()
						messages := plugin.Run(registry.Arguments{Path: fileName, Table: tbl})

						for _, message := range messages {
							if common.IntInArrayInt(cnf.IgnoreCodes(tbl.Name), key+message.Code) {
								continue
							}

							msgComment += fmt.Sprintf("- **[%d]** %s\\n", key+message.Code, message.Message)
						}
					}
				}
				msgComment += "\\n"
			} else {
				msgComment += "Looks ok.\\n"
			}
		}
	})

	if *fComment {
		github.Comment(*fToken, *fRepository, *fPullRequest, msgComment)
	}

	os.Exit(exitCode)
}

func help(rc int) {
	fmt.Printf(USAGE, VERSION)
	os.Exit(rc)
}
