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
	registry.Add("File", func() registry.Method { return &File{} })
}

func (f *File) Run(p registry.Property) {
	f.Property = p
	f.Property.Code = 100

	f.NoEmpty()
	f.WithExtension()
	f.IsUTF8()
	f.EndWithSemicolon()
	f.EndWithNewLine()

	for _, message := range f.Property.Messages {
		fmt.Println(fmt.Sprintf("- [%d] %s", f.Property.Code+message.Code, message.Message))
	}
}

func (f *File) AddMessage(id int, m string) {
	f.Property.Messages = append(f.Property.Messages, registry.Message{Code: id, Message: m})
}

func (f *File) NoEmpty() {
	if len(f.Property.Table.Raw) == 0 {
		f.AddMessage(1, "File is empty.")
	}
}

func (f *File) WithExtension() {
	if filepath.Ext(f.Property.FilePath) != ".sql" {
		f.AddMessage(2, "Invalid file extension, should by '.sql'.")
	}
}

func (f *File) IsUTF8() {
	if !utf8.ValidString(f.Property.Table.Raw) {
		f.AddMessage(3, "Invalid UTF-8 encoding.")
	}
}

func (f *File) EndWithSemicolon() {
	ex := `.*;`
	match, err := regexp.MatchString(ex, f.Property.Table.Raw)
	if match == false || err != nil {
		f.AddMessage(4, "No ending with ';'.")
	}
}

func (f *File) EndWithNewLine() {
	ex := `.*;\n`
	match, err := regexp.MatchString(ex, f.Property.Table.Raw)
	if match == false || err != nil {
		f.AddMessage(5, "No ending with new line.")
	}
}
