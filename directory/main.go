package directory

import (
	"io/ioutil"
	"os"
	"path"
)

func ReadFile(filePath string) (string, error) {
	body, err := ioutil.ReadFile(filePath)
	return string(body), err
}

func ReadDir(p string) (files []string) {
	filesInDir, _ := ioutil.ReadDir(p)
	for _, item := range filesInDir {
		if !item.IsDir() {
			files = append(files, path.Join(p, item.Name()))
		}
	}

	return files
}

func Explore(path string, doFile func(fileName, fileContent string)) {
	if IsDir(path) {
		for _, file := range ReadDir(path) {
			data, _ := ReadFile(file)
			doFile(file, data)
		}
	} else {
		data, _ := ReadFile(path)
		doFile(path, data)
	}
}

func IsDir(path string) bool {
	pathInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return pathInfo.IsDir()
}
