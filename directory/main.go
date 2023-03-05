package directory

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"skeely/common"
	"skeely/common/exec"
)

func ReadFile(filePath string) (string, error) {
	body, err := ioutil.ReadFile(filePath)
	return string(body), err
}

func Explore(path string, git bool, doFile func(fileName, fileContent string)) {
	var gitFiles []string

	if git {
		gitFiles = GitChangedFiles()
		fmt.Println(gitFiles)
	}

	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".sql" {
			fmt.Println(path)

			if git && common.StringInSlice(path, gitFiles) {
				data, _ := ReadFile(path)
				doFile(path, data)
			}
			if !git {
				data, _ := ReadFile(path)
				doFile(path, data)
				fmt.Println("Analizo...")
			}
		}
		return nil
	})
}

func GitChangedFiles() (files []string) {
	stdout, exitcode := exec.Command("git diff --diff-filter='ACMRT' --ignore-submodules=all --name-only FETCH_HEAD")
	fmt.Println("exitcode", exitcode)
	fmt.Println("stdout", stdout)
	if exitcode == 0 {
		lines := strings.Split(stdout, "\n")
		for _, line := range lines {
			if len(line) == 0 {
				continue
			}

			if filepath.Ext(line) != ".sql" {
				continue
			}

			files = append(files, line)
		}
	}

	return files
}
