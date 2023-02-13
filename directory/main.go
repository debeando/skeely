package directory

import (
	"io/ioutil"
	"os"
)

func ReadFile(filePath string) (string, error) {
	body, err := ioutil.ReadFile(filePath)
	return string(body), err
}

func ReadDir(path string) (files []string) {
	filesInDir, _ := ioutil.ReadDir(path)
	for _, file := range filesInDir {
		if ! file.IsDir() {
			files = append(files, path + file.Name())
		}
	}

	return files
}

func Explore(path string, doFile func(content string)) {
	if IsDir(path) {
		for _, file := range ReadDir(path) {
			data, _ := ReadFile(file)
			doFile(data)
		}
	} else {
		data, _ := ReadFile(path)
		doFile(data)
	}
}

func IsDir(path string) bool {
	pathInfo, _ := os.Stat(path)
	return pathInfo.IsDir()
}
