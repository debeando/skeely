package help

import (
	"fmt"
	"os"

	"skeely/command/version"
)

const USAGE = `skeely %s Is a Schema Linter for MySQL, this tool help to identifying
some common and uncommon mistakes on data model.

USAGE:
	skeely [--help | [ --path | --files ] | --version]

OPTIONS:
  --comment               Send summary as comment into GitHub.
  --files                 List of files to lint, separated by space.
  --git                   Auto identifying git changed files, require --path option.
  --github-pull-request   Pull request number.
  --github-repository     Repository path on github.
  --github-token          Token to auth in github.
  --help                  Show this help.
  --ignore                List of codes to ignore for all tables, separated by comma.
  --path                  Path of the directory to start to find *.sql to lint.
  --version               Print version numbers.

EXAMPLES:

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

For more help, plese visit: https://github.com/debeando/skeely
`

func Show(rc int) {
	fmt.Printf(USAGE, version.VERSION)
	os.Exit(rc)
}
