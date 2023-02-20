package definition

import (
	"fmt"
	"regexp"
	"strings"

	"mysql-ddl-lint/registry"
)

type Plugin struct {
	Arguments registry.Arguments
	Messages  []registry.Message
}

func init() {
	registry.Add(300, func() registry.Module { return &Plugin{} })
}

func (p *Plugin) Run(a registry.Arguments) []registry.Message {
	p.Arguments = a

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

	return p.Messages
}

func (p *Plugin) AddMessage(id int, message string) {
	p.Messages = append(p.Messages, registry.Message{Code: id, Message: message})
}

func (p *Plugin) IsTable() {
	ex := `CREATE\sTABLE\s\x60.*\x60\s\([\s\S]*\).*`
	match, err := regexp.MatchString(ex, p.Arguments.Table.Raw)
	if match == false || err != nil {
		p.AddMessage(1, "This is not a table definition.")
	}
}

func (p *Plugin) Engine() {
	if p.Arguments.Table.Engine != "InnoDB" {
		p.AddMessage(2, "Table engine is not InnoDB.")
	}
}

func (p *Plugin) Charset() {
	if !strings.Contains(strings.ToLower(p.Arguments.Table.Charset), "utf8") {
		p.AddMessage(3, "Table charset is not set to use UTF8.")
	}
}

func (p *Plugin) Collate() {
	if !strings.Contains(strings.ToLower(p.Arguments.Table.Collate), "utf8") {
		p.AddMessage(4, "Table collate is not set to use UTF8.")
	}
}

func (p *Plugin) Comment() {
	if len(p.Arguments.Table.Comment) == 0 {
		p.AddMessage(5, "Table no have description.")
	}
}

func (p *Plugin) Name() {
	if len(p.Arguments.Table.Name) == 0 {
		p.AddMessage(6, "Table no have name.")
	}
}

func (p *Plugin) Length() {
	if len(p.Arguments.Table.Name) > 64 {
		p.AddMessage(7, fmt.Sprintf("Table name is large: %s.", p.Arguments.Table.Name))
	}
}

func (p *Plugin) Dots() {
	if strings.Contains(p.Arguments.Table.Name, ".") {
		p.AddMessage(8, fmt.Sprintf("Table name contains dot's in the name: %s.", p.Arguments.Table.Name))
	}
}

func (p *Plugin) StartWithUnderscore() {
	if strings.HasPrefix(p.Arguments.Table.Name, "_") {
		p.AddMessage(9, fmt.Sprintf("Table name start with underscore: %s.", p.Arguments.Table.Name))
	}
}

func (p *Plugin) EndWithTemp() {
	if strings.HasSuffix(p.Arguments.Table.Name, "_tmp") || strings.HasSuffix(p.Arguments.Table.Name, "_temp") {
		p.AddMessage(10, fmt.Sprintf("Table name end with _tmp or _temp: %s.", p.Arguments.Table.Name))
	}
}

func (p *Plugin) LowerCase() {
	for _, r := range p.Arguments.Table.Name {
		if r >= 'A' && r <= 'Z' {
			p.AddMessage(11, fmt.Sprintf("Table name has capital letter: %s.", p.Arguments.Table.Name))
			break
		}
	}
}
