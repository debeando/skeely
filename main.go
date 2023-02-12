// TODO:
// =====
// - Hacer una función regexFind() que retorne el valor y refactorizar para daptar a esto.
//   Así se simplifica la expresion regular que hay para la tabla.
// - Profundisar el parser fields.
// - Profundisar el parser constraints.
// - Hacer un stdout en formato tabla con el resumen de todas las tablas analizadas.

package main

import (
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

	directory.Explore(*fPath, func(content []byte) {
		t := table.Table{}
		if ! t.Parser(content) {
			return
		}

		fmt.Println("======")
		fmt.Println("Table:", t.Name)
		fmt.Printf("Fields: %+v\n", t.Fields)
		fmt.Println("Engine:", t.Engine)
		fmt.Println("Charset:", t.Charset)
		fmt.Println("Collate:", t.Collate)
		fmt.Println("RowFormat:", t.RowFormat)
		fmt.Println("Comment:", t.Comment)
		fmt.Printf("PrimaryKey: %+v\n", t.PrimaryKey)
		fmt.Printf("UniqueKeys: %+v\n", t.UniqueKeys)
		fmt.Printf("Keys: %+v\n", t.Keys)
		fmt.Printf("Constraints: %+v\n", t.Constraints)
	})
}

func help(rc int) {
	fmt.Printf(USAGE, VERSION)
	os.Exit(rc)
}
