package command

import (
	"flag"
	"fmt"
	"os"

	"skeely/common"
	"skeely/common/github"
	"skeely/config"
	"skeely/directory"
	"skeely/registry"
	"skeely/table"

	_ "skeely/plugins"
)

var exitCode = 0

func init() {
	skeelyPATH := os.Getenv("INPUT_PATH")
	skeelyCOMMENT := os.Getenv("INPUT_COMMENT")
	skeelyTOKEN := os.Getenv("INPUT_TOKEN")
	skeelyREPOSITORY := os.Getenv("INPUT_REPOSITORY")
	skeelyPULLREQUEST := os.Getenv("INPUT_PULLREQUEST")

	fmt.Println("PATH:", skeelyPATH)
	fmt.Println("COMMENT:", skeelyCOMMENT)
	fmt.Println("TOKEN:", skeelyTOKEN)
	fmt.Println("REPOSITORY:", skeelyREPOSITORY)
	fmt.Println("PULLREQUEST:", skeelyPULLREQUEST)

	os.Exit(0)

	fGitHubComment := flag.Bool("github-comment", false, "")
	fGitHubPullRequest := flag.Int("github-pull-request", 0, "")
	fGitHubRepository := flag.String("github-repository", "", "")
	fGitHubToken := flag.String("github-token", "", "")
	fHelp := flag.Bool("help", false, "")
	fPath := flag.String("path", "", "")
	fVersion := flag.Bool("version", false, "")
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
	msgComment = "# MySQL Migration linter summary\\n"

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

		if *fGitHubComment {
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

	if *fGitHubComment {
		github.Comment(*fGitHubToken, *fGitHubRepository, *fGitHubPullRequest, msgComment)
	}

	os.Exit(exitCode)
}

func Run() {
}
