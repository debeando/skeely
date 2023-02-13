package table

type Table struct {
	Charset string
	Collate string
	Comment string
	Constraints []Constraint
	Engine string
	Fields []Field
	Keys []Key
	Name string
	PrimaryKey []string
	Raw string `json:"-"`
	RowFormat string
	UniqueKeys []Key
}

func (t *Table) Parser(r string) error {
	(*t) = Table { Raw: r }

	t.GetCharset()
	t.GetCollate()
	t.GetComment()
	t.GetConstraints()
	t.GetEngine()
	t.GetFields()
	t.GetKeys()
	t.GetName()
	t.GetPrimaryKey()
	t.GetRowFormat()
	t.GetUniqueKeys()

	return nil
}

func (t *Table) GetName() {
	ex := `^CREATE TABLE \x60(?P<Table>[0-9,a-z,A-Z$_\.]+)\x60\s\(`
	t.Name = findMatchOne(ex, t.Raw)
}

func (t *Table) GetEngine() {
	ex := `ENGINE=(?P<Engine>\w+)`
	t.Engine = findMatchOne(ex, t.Raw)
}

func (t *Table) GetCharset() {
	ex := `DEFAULT CHARSET=(?P<Charset>\w+)`
	t.Charset = findMatchOne(ex, t.Raw)
}

func (t *Table) GetCollate() {
	ex := `COLLATE=(?P<Collate>\w+)`
	t.Collate = findMatchOne(ex, t.Raw)
}

func (t *Table) GetRowFormat() {
	ex := `ROW_FORMAT=(?P<RowFormat>\w+)`
	t.RowFormat = findMatchOne(ex, t.Raw)
}

func (t *Table) GetComment() {
	ex := `COMMENT='(?P<Comment>.+)'`
	t.Comment = findMatchOne(ex, t.Raw)
}

func (t *Table) GetPrimaryKey() {
	ex := `\s{2}PRIMARY KEY \((?P<Fields>(\x60.+\x60(, )?)+)\)`
	t.PrimaryKey = stringToArray(findMatchOne(ex, t.Raw))
}

func (t *Table) GetUniqueKeys() {
	ex := `\s{2}UNIQUE KEY \x60(?P<Name>[0-9,a-z,A-Z$_\.]+)\x60 \((?P<Fields>(\x60.+\x60(, )?)+)\)`
	for _, item := range find(ex, t.Raw) {
		t.UniqueKeys = append(t.UniqueKeys, Key{
			Name: item[1],
			Fields: stringToArray(item[2]),
		})
	}
}

func (t *Table) GetKeys() {
	ex := `\s{2}KEY \x60(?P<Name>[0-9,a-z,A-Z$_\.]+)\x60 \((?P<Fields>(\x60.+\x60(, )?)+)\)`
	for _, item := range find(ex, t.Raw) {
		t.Keys = append(t.Keys, Key{
			Name: item[1],
			Fields: stringToArray(item[2]),
		})
	}
}

func (t *Table) GetConstraints() {
	ex := `\s{2}CONSTRAINT \x60(?P<Name>[0-9,a-z,A-Z$_\.]+)\x60\s(?P<Properties>.*)`
	for _, item := range find(ex, t.Raw) {
		t.Constraints = append(t.Constraints, Constraint{
			Name: item[1],
		})
	}
}

func (t *Table) GetFields() {
	ex := `\s{2}(?P<Field>\x60.*)`
	for _, item := range find(ex, t.Raw) {
		if len(item) > 0 {
			f := Field{}
			f.Parser(item[0])
			t.Fields = append(t.Fields, f)
		}
	}
}
