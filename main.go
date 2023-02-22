package main

import (
	"flag"
	"fmt"
	"os"

	"mylinter/common"
	"mylinter/config"
	"mylinter/directory"
	"mylinter/registry"
	"mylinter/table"

	_ "mylinter/plugins"
)

const VERSION string = "0.0.0-beta.1"
const USAGE = `mylinter %s.`

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

	cnf := config.Config{}
	err := cnf.Load()
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

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
