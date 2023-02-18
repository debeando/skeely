package main

import (
	"flag"
	"fmt"
	"os"

	"mysql-ddl-lint/directory"
	"mysql-ddl-lint/plugins/registry"
	"mysql-ddl-lint/table"

	_ "mysql-ddl-lint/plugins"
)

const VERSION string = "1.0.0"
const USAGE = `mysql-ddl-lint %s.`

var exitCode = 0

func main() {
	fHelp := flag.Bool("help", false, "")
	fVersion := flag.Bool("version", false, "")
	fPath := flag.String("path", "", "")
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

	directory.Explore(*fPath, func(fileName, fileContent string) {
		fmt.Println("> File:", fileName)

		t := table.Table{}
		if t.Parser(fileContent) != nil {
			return
		}

		for key := range registry.Plugins {
			if creator, ok := registry.Plugins[key]; ok {
				properties := creator().Run(registry.Property {
					FilePath: fileName,
					Table:    t,
				})

				for _, message := range properties.Messages {
					fmt.Println(fmt.Sprintf("- [%d] %s", properties.Code + message.Code, message.Message))
					exitCode = 1
				}
			}
		}
	})

	os.Exit(exitCode)
}

func help(rc int) {
	fmt.Printf(USAGE, VERSION)
	os.Exit(rc)
}
