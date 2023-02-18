package definition

import (
	"fmt"
	"regexp"
	"strings"

	"mysql-ddl-lint/plugins/registry"
)

type Plugin struct {
	Property registry.Property
}

func init() {
	registry.Add("Table Definition", func() registry.Method { return &Plugin{} })
}

func (p *Plugin) Run(a registry.Property) registry.Property {
	p.Property = a
	p.Property.Code = 300

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

	return p.Property
}

func (p *Plugin) AddMessage(id int, message string) {
	p.Property.Messages = append(p.Property.Messages, registry.Message{Code: id, Message: message})
}

func (p *Plugin) IsTable() {
	ex := `CREATE\sTABLE\s\x60.*\x60\s\([\s\S]*\).*`
	match, err := regexp.MatchString(ex, p.Property.Table.Raw)
	if match == false || err != nil {
		p.AddMessage(1, "This is not a table definition.")
	}
}

func (p *Plugin) Engine() {
	if p.Property.Table.Engine != "InnoDB" {
		p.AddMessage(2, "Table engine is not InnoDB.")
	}
}

func (p *Plugin) Charset() {
	if !strings.Contains(strings.ToLower(p.Property.Table.Charset), "utf8") {
		p.AddMessage(3, "Table charset is not set to use UTF8.")
	}
}

func (p *Plugin) Collate() {
	if !strings.Contains(strings.ToLower(p.Property.Table.Collate), "utf8") {
		p.AddMessage(4, "Table collate is not set to use UTF8.")
	}
}

func (p *Plugin) Comment() {
	if len(p.Property.Table.Comment) == 0 {
		p.AddMessage(5, "Table no have description.")
	}
}

func (p *Plugin) Name() {
	if len(p.Property.Table.Name) == 0 {
		p.AddMessage(6, "Table no have name.")
	}
}

func (p *Plugin) Length() {
	if len(p.Property.Table.Name) > 64 {
		p.AddMessage(7, fmt.Sprintf("Table name is large: %s.", p.Property.Table.Name))
	}
}

func (p *Plugin) Dots() {
	if strings.Contains(p.Property.Table.Name, ".") {
		p.AddMessage(8, fmt.Sprintf("Table name contains dot's in the name: %s.", p.Property.Table.Name))
	}
}

func (p *Plugin) StartWithUnderscore() {
	if strings.HasPrefix(p.Property.Table.Name, "_") {
		p.AddMessage(9, fmt.Sprintf("Table name start with underscore: %s.", p.Property.Table.Name))
	}
}

func (p *Plugin) EndWithTemp() {
	if strings.HasSuffix(p.Property.Table.Name, "_tmp") || strings.HasSuffix(p.Property.Table.Name, "_temp") {
		p.AddMessage(10, fmt.Sprintf("Table name end with _tmp or _temp: %s.", p.Property.Table.Name))
	}
}

func (p *Plugin) LowerCase() {
	for _, r := range p.Property.Table.Name {
		if r >= 'A' && r <= 'Z' {
			p.AddMessage(11, fmt.Sprintf("Table name has capital letter: %s.", p.Property.Table.Name))
			break
		}
	}
}
