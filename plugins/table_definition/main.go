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
	registry.Add("Table Definition", func() registry.Method{ return &TableDefinition{} })
}

func (m *TableDefinition) Run(p registry.Property) {
	m.Property = p

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
}

func (m *TableDefinition) IsTable() {
    ex := `CREATE\sTABLE\s\x60.*\x60\s\([\s\S]*\).*`
    match, err := regexp.MatchString(ex, m.Property.Table.Raw)
    if match == false || err != nil {
        fmt.Println("[301] This is not a table definition.")
    }
}

func (m *TableDefinition) Engine() {
    if m.Property.Table.Engine != "InnoDB" {
        fmt.Println("[302] Table engine is not InnoDB.")
    }
}

func (m *TableDefinition) Charset() {
    if ! strings.Contains(strings.ToLower(m.Property.Table.Charset), "utf8") {
        fmt.Println("[303] Table charset is not set to use UTF8.")
    }
}

func (m *TableDefinition) Collate() {
    if ! strings.Contains(strings.ToLower(m.Property.Table.Collate), "utf8") {
        fmt.Println("[304] Table collate is not set to use UTF8.")
    }
}

func (m *TableDefinition) Comment() {
    if len(m.Property.Table.Comment) == 0 {
        fmt.Println("[305] Table no have description.")
    }
}

func (m *TableDefinition) Name() {
    if len(m.Property.Table.Name) == 0 {
        fmt.Println("[306] Table no have name.")
    }
}

func (m *TableDefinition) Length() {
    if len(m.Property.Table.Name) > 64 {
        fmt.Println("[307] Table name is large.")
    }
}

func (m *TableDefinition) Dots() {
    if strings.Contains(m.Property.Table.Name, ".") {
        fmt.Println("[308] Table name contains dot's in the name.") 
    }
}

func (m *TableDefinition) StartWithUnderscore() {
    if strings.HasPrefix(m.Property.Table.Name, "_") {
        fmt.Println("[309] Table name start with underscore.")  
    }
}

func (m *TableDefinition) EndWithTemp() {
    if strings.HasSuffix(m.Property.Table.Name, "_tmp") {
        fmt.Println("[310] Table name end with _tmp.")  
    }
    if strings.HasSuffix(m.Property.Table.Name, "_temp") {
        fmt.Println("[310] Table name end with _temp.") 
    }
}

func (m *TableDefinition) LowerCase() {
    for _, r := range m.Property.Table.Name {
        if (r >= 'A' && r <= 'Z') {
            fmt.Println("[311] Table name has capital letter.")
            break
        }
    }
}
