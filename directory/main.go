package directory

import (
	// "fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"skeely/flags"
	// "skeely/common"
)

func ReadFile(filePath string) (string, error) {
	body, err := ioutil.ReadFile(filePath)
	return string(body), err
}

func Explore(doFile func(fileName, fileContent string)) {
	f := flags.GetInstance()

	var files []string
	// e, _ := exists(path)
	// // fmt.Println(e, err)

	// if e == false {
	// 	files_temp := strings.Split(path, " ")
	// 	for _, file := range files_temp {
	// 		e, _ = exists(file)
	// 		if e {
	// 			files = append(files, file)
	// 		}
	// 		// fmt.Println(e)
	// 	}

	// 	path = "."
	// }

	// if git {
	// 	gitFiles = GitChangedFiles()
	// 	fmt.Println(gitFiles)
	// }

	if exists(f.Path) {
		filepath.Walk(f.Path, func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".sql" {
				// fmt.Println(path)
				// if common.StringInSlice(path, files) {
				// 	fmt.Println("match...")
				// // 	data, _ := ReadFile(path)
				// // 	doFile(path, data)
				// }
				// if !git {
				data, _ := ReadFile(path)
				doFile(path, data)
				// fmt.Println("Analizo...")
				// }
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

// TODO: cambiar el nombre de directory to algo....
// TODO: poner un nuevo arg que sea files
