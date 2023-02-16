package table_definition

import (
	"fmt"
	"regexp"
	"strings"

	"mysql-ddl-lint/plugins/registry"
)

type TableDefinition struct {
	Property registry.Property
}

func init() {
	registry.Add("Table Definition", func() registry.Method { return &TableDefinition{} })
}

func (m *TableDefinition) Run(p registry.Property) {
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
		fmt.Println(fmt.Sprintf("[%d] %s", m.Property.Code+message.Code, message.Message))
	}
}

func (m *TableDefinition) AddMessage(id int, message string) {
	m.Property.Messages = append(m.Property.Messages, registry.Message{Code: id, Message: message})
}

func (m *TableDefinition) IsTable() {
	ex := `CREATE\sTABLE\s\x60.*\x60\s\([\s\S]*\).*`
	match, err := regexp.MatchString(ex, m.Property.Table.Raw)
	if match == false || err != nil {
		m.AddMessage(1, "This is not a table definition.")
	}
}

func (m *TableDefinition) Engine() {
	if m.Property.Table.Engine != "InnoDB" {
		m.AddMessage(2, "Table engine is not InnoDB.")
	}
}

func (m *TableDefinition) Charset() {
	if !strings.Contains(strings.ToLower(m.Property.Table.Charset), "utf8") {
		m.AddMessage(3, "Table charset is not set to use UTF8.")
	}
}

func (m *TableDefinition) Collate() {
	if !strings.Contains(strings.ToLower(m.Property.Table.Collate), "utf8") {
		m.AddMessage(4, "Table collate is not set to use UTF8.")
	}
}

func (m *TableDefinition) Comment() {
	if len(m.Property.Table.Comment) == 0 {
		m.AddMessage(5, "Table no have description.")
	}
}

func (m *TableDefinition) Name() {
	if len(m.Property.Table.Name) == 0 {
		m.AddMessage(6, "Table no have name.")
	}
}

func (m *TableDefinition) Length() {
	if len(m.Property.Table.Name) > 64 {
		m.AddMessage(7, "Table name is large.")
	}
}

func (m *TableDefinition) Dots() {
	if strings.Contains(m.Property.Table.Name, ".") {
		m.AddMessage(8, "Table name contains dot's in the name.")
	}
}

func (m *TableDefinition) StartWithUnderscore() {
	if strings.HasPrefix(m.Property.Table.Name, "_") {
		m.AddMessage(9, "Table name start with underscore.")
	}
}

func (m *TableDefinition) EndWithTemp() {
	if strings.HasSuffix(m.Property.Table.Name, "_tmp") || strings.HasSuffix(m.Property.Table.Name, "_temp") {
		m.AddMessage(10, "Table name end with _tmp or _temp.")
	}
}

func (m *TableDefinition) LowerCase() {
	for _, r := range m.Property.Table.Name {
		if r >= 'A' && r <= 'Z' {
			m.AddMessage(11, "Table name has capital letter.")
			break
		}
	}
}
