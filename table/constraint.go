package table

import (
	"mysql-ddl-lint/common"
)

type Constraint struct {
	Delete          bool
	ForeignKeys     []string
	Name            string
	Raw             string `json:"-"`
	ReferenceTable  string
	ReferenceFields []string
	Update          bool
}

func (c *Constraint) Parser(r string) error {
	(*c) = Constraint{Raw: r}

	c.GetDelete()
	c.GetName()
	c.GetUpdate()
	c.GetForeignKeys()
	c.GetReferenceTable()
	c.GetReferenceFields()

	return nil
}

func (c *Constraint) GetName() {
	ex := `\s{2}CONSTRAINT \x60(?P<Name>[0-9,a-z,A-Z$_\.]+)\x60\s`
	c.Name = common.FindMatchOne(ex, c.Raw, 1)
}

func (c *Constraint) GetDelete() {
	c.Delete = common.StringIn(c.Raw, " DELETE")
}

func (c *Constraint) GetUpdate() {
	c.Update = common.StringIn(c.Raw, " UPDATE")
}

func (c *Constraint) GetForeignKeys() {
	ex := `\s{2}CONSTRAINT\s\x60[0-9,a-z,A-Z$_\.]+\x60\sFOREIGN KEY\s(?:\((?P<ForeignKeys>\x60.+\x60)\))\sREFERENCES.*`
	c.ForeignKeys = common.StringToArray(common.FindMatchOne(ex, c.Raw, 1))
}

func (c *Constraint) GetReferenceTable() {
	ex := `\sREFERENCES\s\x60(?P<Table>[0-9,a-z,A-Z$_\.]+)?\x60\s`
	c.ReferenceTable = common.FindMatchOne(ex, c.Raw, 1)
}

func (c *Constraint) GetReferenceFields() {
	ex := `\sREFERENCES\s\x60.*\s\((?P<Fields>\x60.*\x60)\)`
	c.ReferenceFields = common.StringToArray(common.FindMatchOne(ex, c.Raw, 1))
}
