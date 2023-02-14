package file

import (
	"fmt"
	"path/filepath"
	"regexp"
	"unicode/utf8"

	"mysql-ddl-lint/plugins/registry"
)

type File struct {
	Property registry.Property
}

func init() {
	registry.Add("File", func() registry.Method{ return &File{} })
}

func (f *File) Run(p registry.Property) {
	f.Property = p

	f.NoEmpty()
	f.WithExtension()
	f.IsUTF8()
	f.EndWithSemicolon()
	f.EndWithNewLine()
}

func (f *File) NoEmpty() {
    if len(f.Property.Table.Raw) == 0 {
        fmt.Println("[101] File is empty.")
    }
}

func (f *File) WithExtension() {
	if filepath.Ext(f.Property.FilePath) != ".sql" {
		fmt.Println("[102] Invalid file extension, should by '.sql'.")
	}
}

func (f *File) IsUTF8() {
	if !utf8.ValidString(f.Property.Table.Raw) {
		fmt.Println("[103] Invalid UTF-8 encoding.")
	}
}

func (f *File) EndWithSemicolon() {
	ex := `.*;`
	match, err := regexp.MatchString(ex, f.Property.Table.Raw)
	if match == false || err != nil {
		fmt.Println("[104] No ending with ';'.")
	}
}

func (f *File) EndWithNewLine() {
	ex := `.*;\n`
	match, err := regexp.MatchString(ex, f.Property.Table.Raw)
	if match == false || err != nil {
		fmt.Println("[105] No ending with new line.")
	}
}
