package definition

import (
	"fmt"
	"regexp"
	"strings"

	"skeely/message"
	"skeely/registry"
)

type Plugin struct {
	Arguments registry.Arguments
	Messages  message.Messages
	Incidents message.Plugin
}

func init() {
	registry.Add(300, func() registry.Module { return &Plugin{} })
}

func (p *Plugin) Run(a registry.Arguments) message.Plugin {
	p.Arguments = a
	p.Messages = message.Messages{
		{Code: 1, Message: "This is not a table definition."},
		{Code: 2, Message: "Table engine is not InnoDB."},
		{Code: 3, Message: "Table charset is not set to use UTF8."},
		{Code: 4, Message: "Table collate is not set to use UTF8."},
		{Code: 5, Message: "Table no have description."},
		{Code: 6, Message: "Table no have name."},
		{Code: 7, Message: "Table name is large: %s"},
		{Code: 8, Message: "Table name contains dot's in the name: %s"},
		{Code: 9, Message: "Table name start with underscore: %s"},
		{Code: 10, Message: "Table name end with _tmp or _temp: %s"},
		{Code: 11, Message: "Table name has capital letter: %s"},
	}

	p.IsTable()
	p.Engine()
	p.Charset()
	p.Collate()
	p.Comment()
	p.Name()
	p.Length()
	p.Dots()
	p.StartWithUnderscore()
	p.EndWithTemp()
	p.LowerCase()

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

func (p *Plugin) IsTable() {
	ex := `CREATE\sTABLE\s\x60.*\x60\s*\([\s\S]*\).*`
	match, err := regexp.MatchString(ex, p.Arguments.Table.Raw)
	if match == false || err != nil {
		p.AddMessage(1)
	}
}

func (p *Plugin) Engine() {
	if p.Arguments.Table.Engine != "InnoDB" {
		p.AddMessage(2)
	}
}

func (p *Plugin) Charset() {
	if !strings.Contains(strings.ToLower(p.Arguments.Table.Charset), "utf8") {
		p.AddMessage(3)
	}
}

func (p *Plugin) Collate() {
	if !strings.Contains(strings.ToLower(p.Arguments.Table.Collate), "utf8") {
		p.AddMessage(4)
	}
}

func (p *Plugin) Comment() {
	if len(p.Arguments.Table.Comment) == 0 {
		p.AddMessage(5)
	}
}

func (p *Plugin) Name() {
	if len(p.Arguments.Table.Name) == 0 {
		p.AddMessage(6)
	}
}

func (p *Plugin) Length() {
	// https://dev.mysql.com/doc/refman/8.0/en/identifier-length.html
	if len(p.Arguments.Table.Name) > 64 {
		p.AddMessage(7, p.Arguments.Table.Name)
	}
}

func (p *Plugin) Dots() {
	if strings.Contains(p.Arguments.Table.Name, ".") {
		p.AddMessage(8, p.Arguments.Table.Name)
	}
}

func (p *Plugin) StartWithUnderscore() {
	if strings.HasPrefix(p.Arguments.Table.Name, "_") {
		p.AddMessage(9, p.Arguments.Table.Name)
	}
}

func (p *Plugin) EndWithTemp() {
	if strings.HasSuffix(p.Arguments.Table.Name, "_tmp") || strings.HasSuffix(p.Arguments.Table.Name, "_temp") {
		p.AddMessage(10, p.Arguments.Table.Name)
	}
}

func (p *Plugin) LowerCase() {
	for _, r := range p.Arguments.Table.Name {
		if r >= 'A' && r <= 'Z' {
			p.AddMessage(11, p.Arguments.Table.Name)
			break
		}
	}
}
