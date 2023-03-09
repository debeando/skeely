package directory

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"skeely/flags"
)

func ReadFile(filePath string) (string, error) {
	body, err := ioutil.ReadFile(filePath)
	return string(body), err
}

func Explore(doFile func(fileName, fileContent string)) {
	f := flags.GetInstance()

	var files []string

	if exists(f.Path) {
		filepath.Walk(f.Path, func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".sql" {
				data, _ := ReadFile(path)
				doFile(path, data)
			}
			return nil
		})
	}

	if !exists(f.Path) && len(f.Files) > 0 {
		files_temp := strings.Split(f.Files, " ")
		if len(files_temp) == 0 {
			return
		}
		for _, file := range files_temp {
			if exists(file) {
				files = append(files, file)
			}
		}

		for _, file := range files {
			data, _ := ReadFile(file)
			doFile(file, data)
		}
	}
}

func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
