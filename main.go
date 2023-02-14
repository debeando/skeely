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
		fmt.Println("--> Lint file:", fileName)

		t := table.Table{}
		if t.Parser(fileContent) != nil {
			return
		}

		for key := range registry.Plugins {
			if creator, ok := registry.Plugins[key]; ok {
				plugin := creator()
				plugin.Run(registry.Property{
					FilePath: fileName,
					Table:    t,
				})
			}
		}
	})
}

func help(rc int) {
	fmt.Printf(USAGE, VERSION)
	os.Exit(rc)
}
