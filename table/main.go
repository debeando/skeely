package table

type Table struct {
	Raw []byte
	Name string
	Fields []Field
	Engine string
	Charset string
	Collate string
	RowFormat string
	Comment string
	PrimaryKey []string
	UniqueKeys []Key
	Keys []Key
	Constraints []Constraint
}

type Field struct {
	Name string
	Type string
	Null bool
	Default string
	Comment string
}

type Key struct {
	Name string
	Fields []string
}

type Constraint struct {
	Name string
	Fields []string
	Reference []Key
	Delete string
	Update string
}
