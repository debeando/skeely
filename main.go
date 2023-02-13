package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"mysql-ddl-checker/directory"
	"mysql-ddl-checker/table"
)

const VERSION string = "1.0.0"
const USAGE = `mysql-ddl-checker %s.`

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

	directory.Explore(*fPath, func(file string) {
		t := table.Table{}
		if t.Parser(file) != nil {
			return
		}

		out, _ := json.Marshal(t)
		fmt.Printf(string(out))
	})
}

func help(rc int) {
	fmt.Printf(USAGE, VERSION)
	os.Exit(rc)
}
