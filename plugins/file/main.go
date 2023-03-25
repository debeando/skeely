package file

import (
	"fmt"
	"regexp"
	"unicode/utf8"

	"skeely/message"
	"skeely/registry"
)

type Plugin struct {
	Arguments registry.Arguments
	Messages  message.Messages
	Incidents message.Plugin
}

func init() {
	registry.Add(100, func() registry.Module { return &Plugin{} })
}

func (p *Plugin) Run(a registry.Arguments) message.Plugin {
	p.Arguments = a
	p.Messages = []message.Message{
		{Code: 1, Message: "File is empty."},
		{Code: 2, Message: "Invalid UTF-8 encoding."},
		{Code: 3, Message: "No ending with ';'."},
		{Code: 4, Message: "No ending with new line."},
		// TODO: Code 5: You have many tables on single file.
	}

	p.NoEmpty()
	p.IsUTF8()
	p.EndWithSemicolon()
	p.EndWithNewLine()

	return p.Incidents
}

func (p *Plugin) GetMessage(id int) string {
	for _, message := range p.Messages {
		if message.Code == id {
			return message.Message
		}
	}
	return ""
}

func (p *Plugin) AddMessage(id int, vals ...any) {
	p.Incidents.Add(message.Message{
		Code:    id,
		Message: fmt.Sprintf(p.GetMessage(id), vals...),
	})
}

func (p *Plugin) NoEmpty() {
	if len(p.Arguments.Table.Raw) == 0 {
		p.AddMessage(1)
	}
}

func (p *Plugin) IsUTF8() {
	if !utf8.ValidString(p.Arguments.Table.Raw) {
		p.AddMessage(2)
	}
}

func (p *Plugin) EndWithSemicolon() {
	ex := `.*;`
	match, err := regexp.MatchString(ex, p.Arguments.Table.Raw)
	if match == false || err != nil {
		p.AddMessage(3)
	}
}

func (p *Plugin) EndWithNewLine() {
	ex := `.*;\n`
	match, err := regexp.MatchString(ex, p.Arguments.Table.Raw)
	if match == false || err != nil {
		p.AddMessage(4)
	}
}
