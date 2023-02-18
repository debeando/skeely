package file

import (
	"path/filepath"
	"regexp"
	"unicode/utf8"

	"mysql-ddl-lint/plugins/registry"
)

type Plugin struct {
	Property registry.Property
}

func init() {
	registry.Add("File", func() registry.Method { return &Plugin{} })
}

func (p *Plugin) Run(a registry.Property) registry.Property {
	p.Property = a
	p.Property.Code = 100

	p.NoEmpty()
	p.WithExtension()
	p.IsUTF8()
	p.EndWithSemicolon()
	p.EndWithNewLine()

	return p.Property
}

func (p *Plugin) AddMessage(id int, m string) {
	p.Property.Messages = append(p.Property.Messages, registry.Message{Code: id, Message: m})
}

func (p *Plugin) NoEmpty() {
	if len(p.Property.Table.Raw) == 0 {
		p.AddMessage(1, "File is empty.")
	}
}

func (p *Plugin) WithExtension() {
	if filepath.Ext(p.Property.FilePath) != ".sql" {
		p.AddMessage(2, "Invalid file extension, should by '.sql'.")
	}
}

func (p *Plugin) IsUTF8() {
	if !utf8.ValidString(p.Property.Table.Raw) {
		p.AddMessage(3, "Invalid UTF-8 encoding.")
	}
}

func (p *Plugin) EndWithSemicolon() {
	ex := `.*;`
	match, err := regexp.MatchString(ex, p.Property.Table.Raw)
	if match == false || err != nil {
		p.AddMessage(4, "No ending with ';'.")
	}
}

func (p *Plugin) EndWithNewLine() {
	ex := `.*;\n`
	match, err := regexp.MatchString(ex, p.Property.Table.Raw)
	if match == false || err != nil {
		p.AddMessage(5, "No ending with new line.")
	}
}
