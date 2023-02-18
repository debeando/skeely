package definition

import (
	"fmt"
	"regexp"
	"strings"

	"mysql-ddl-lint/plugins/registry"
)

type Definition struct {
	Property registry.Property
}

func init() {
	registry.Add("Table Definition", func() registry.Method { return &Definition{} })
}

func (m *Definition) Run(p registry.Property) {
	m.Property = p
	m.Property.Code = 300

	m.IsTable()
	m.Engine()
	m.Charset()
	m.Collate()
	m.Comment()
	m.Name()
	m.Length()
	m.Dots()
	m.StartWithUnderscore()
	m.EndWithTemp()
	m.LowerCase()

	for _, message := range m.Property.Messages {
		fmt.Println(fmt.Sprintf("- [%d] %s", m.Property.Code+message.Code, message.Message))
	}
}

func (m *Definition) AddMessage(id int, message string) {
	m.Property.Messages = append(m.Property.Messages, registry.Message{Code: id, Message: message})
}

func (m *Definition) IsTable() {
	ex := `CREATE\sTABLE\s\x60.*\x60\s\([\s\S]*\).*`
	match, err := regexp.MatchString(ex, m.Property.Table.Raw)
	if match == false || err != nil {
		m.AddMessage(1, "This is not a table definition.")
	}
}

func (m *Definition) Engine() {
	if m.Property.Table.Engine != "InnoDB" {
		m.AddMessage(2, "Table engine is not InnoDB.")
	}
}

func (m *Definition) Charset() {
	if !strings.Contains(strings.ToLower(m.Property.Table.Charset), "utf8") {
		m.AddMessage(3, "Table charset is not set to use UTF8.")
	}
}

func (m *Definition) Collate() {
	if !strings.Contains(strings.ToLower(m.Property.Table.Collate), "utf8") {
		m.AddMessage(4, "Table collate is not set to use UTF8.")
	}
}

func (m *Definition) Comment() {
	if len(m.Property.Table.Comment) == 0 {
		m.AddMessage(5, "Table no have description.")
	}
}

func (m *Definition) Name() {
	if len(m.Property.Table.Name) == 0 {
		m.AddMessage(6, "Table no have name.")
	}
}

func (m *Definition) Length() {
	if len(m.Property.Table.Name) > 64 {
		m.AddMessage(7, fmt.Sprintf("Table name is large: %s.", m.Property.Table.Name))
	}
}

func (m *Definition) Dots() {
	if strings.Contains(m.Property.Table.Name, ".") {
		m.AddMessage(8, fmt.Sprintf("Table name contains dot's in the name: %s.", m.Property.Table.Name))
	}
}

func (m *Definition) StartWithUnderscore() {
	if strings.HasPrefix(m.Property.Table.Name, "_") {
		m.AddMessage(9, fmt.Sprintf("Table name start with underscore: %s.", m.Property.Table.Name))
	}
}

func (m *Definition) EndWithTemp() {
	if strings.HasSuffix(m.Property.Table.Name, "_tmp") || strings.HasSuffix(m.Property.Table.Name, "_temp") {
		m.AddMessage(10, fmt.Sprintf("Table name end with _tmp or _temp: %s.", m.Property.Table.Name))
	}
}

func (m *Definition) LowerCase() {
	for _, r := range m.Property.Table.Name {
		if r >= 'A' && r <= 'Z' {
			m.AddMessage(11, fmt.Sprintf("Table name has capital letter: %s.", m.Property.Table.Name))
			break
		}
	}
}
