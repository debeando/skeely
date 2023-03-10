package table

import (
	"skeely/common"
)

type Table struct {
	Charset     string
	Collate     string
	Comment     string
	Constraints []Constraint
	Engine      string
	Fields      []Field
	Keys        []Key
	Name        string
	PrimaryKey  []string
	Raw         string `json:"-"`
	RowFormat   string
	UniqueKeys  []Key
}

func (t *Table) Parser(r string) error {
	(*t) = Table{Raw: r}

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
	ex := `^CREATE TABLE \x60(?P<Table>[0-9,a-z,A-Z$_\.]+)\x60\s*\(`
	t.Name = common.FindMatchOne(ex, t.Raw, 1)
}

func (t *Table) GetEngine() {
	ex := `\sENGINE=(?P<Engine>\w+)`
	t.Engine = common.FindMatchOne(ex, t.Raw, 1)
}

func (t *Table) GetCharset() {
	ex := `\sDEFAULT CHARSET=(?P<Charset>\w+)`
	t.Charset = common.FindMatchOne(ex, t.Raw, 1)
}

func (t *Table) GetCollate() {
	ex := `\sCOLLATE=(?P<Collate>\w+)`
	t.Collate = common.FindMatchOne(ex, t.Raw, 1)
}

func (t *Table) GetRowFormat() {
	ex := `\sROW_FORMAT=(?P<RowFormat>\w+)`
	t.RowFormat = common.FindMatchOne(ex, t.Raw, 1)
}

func (t *Table) GetComment() {
	ex := `\sCOMMENT='(?P<Comment>.+)'`
	t.Comment = common.FindMatchOne(ex, t.Raw, 1)
}

func (t *Table) GetPrimaryKey() {
	ex := `\s{2}PRIMARY KEY\s*\((?P<Fields>(\x60.+\x60(, )?)+)\)`
	t.PrimaryKey = common.StringToArray(common.FindMatchOne(ex, t.Raw, 1))
}

func (t *Table) GetUniqueKeys() {
	ex := `\s{2}UNIQUE\sKEY.*`
	for _, item := range common.Find(ex, t.Raw) {
		if len(item) > 0 {
			c := Key{}
			c.Parser(item[0])
			t.Keys = append(t.Keys, c)
		}
	}
}

func (t *Table) GetKeys() {
	ex := `\s{2}KEY\s\x60.*`
	for _, item := range common.Find(ex, t.Raw) {
		if len(item) > 0 {
			c := Key{}
			c.Parser(item[0])
			t.Keys = append(t.Keys, c)
		}
	}
}

func (t *Table) GetConstraints() {
	ex := `\s{2}(?P<Constraint>CONSTRAINT.*)`
	for _, item := range common.Find(ex, t.Raw) {
		if len(item) > 0 {
			c := Constraint{}
			c.Parser(item[0])
			t.Constraints = append(t.Constraints, c)
		}
	}
}

func (t *Table) GetFields() {
	ex := `\s{2}(?P<Field>\x60.*)`
	for _, item := range common.Find(ex, t.Raw) {
		if len(item) > 0 {
			f := Field{}
			f.Parser(item[0])
			t.Fields = append(t.Fields, f)
		}
	}
}
