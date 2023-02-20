package file

import (
	"path/filepath"
	"regexp"
	"unicode/utf8"

	"mysql-ddl-lint/registry"
)

type Plugin struct {
	Arguments registry.Arguments
	Messages  []registry.Message
}

func init() {
	registry.Add(100, func() registry.Module { return &Plugin{} })
}

func (p *Plugin) Run(a registry.Arguments) []registry.Message {
	p.Arguments = a

	p.NoEmpty()
	p.WithExtension()
	p.IsUTF8()
	p.EndWithSemicolon()
	p.EndWithNewLine()

	return p.Messages
}

func (p *Plugin) AddMessage(id int, m string) {
	p.Messages = append(p.Messages, registry.Message{Code: id, Message: m})
}

func (p *Plugin) NoEmpty() {
	if len(p.Arguments.Table.Raw) == 0 {
		p.AddMessage(1, "File is empty.")
	}
}

func (p *Plugin) WithExtension() {
	if filepath.Ext(p.Arguments.Path) != ".sql" {
		p.AddMessage(2, "Invalid file extension, should by '.sql'.")
	}
}

func (p *Plugin) IsUTF8() {
	if !utf8.ValidString(p.Arguments.Table.Raw) {
		p.AddMessage(3, "Invalid UTF-8 encoding.")
	}
}

func (p *Plugin) EndWithSemicolon() {
	ex := `.*;`
	match, err := regexp.MatchString(ex, p.Arguments.Table.Raw)
	if match == false || err != nil {
		p.AddMessage(4, "No ending with ';'.")
	}
}

func (p *Plugin) EndWithNewLine() {
	ex := `.*;\n`
	match, err := regexp.MatchString(ex, p.Arguments.Table.Raw)
	if match == false || err != nil {
		p.AddMessage(5, "No ending with new line.")
	}
}
