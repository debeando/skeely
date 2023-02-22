package table

import (
	"strings"

	"mylinter/common"
)

type Field struct {
	AutoIncrement bool
	Collate       string
	Comment       string
	Default       string
	Enum          []string
	Length        int
	Name          string
	NotNull       bool
	Raw           string `json:"-"`
	Type          string
	Unsigned      bool
	ZeroFill      bool
}

func (f *Field) Parser(r string) error {
	(*f) = Field{Raw: r}

	f.GetAutoIncrement()
	f.GetCollate()
	f.GetComment()
	f.GetDefault()
	f.GetEnum()
	f.GetName()
	f.GetLength()
	f.GetNotNull()
	f.GetType()
	f.GetUnsigned()
	f.GetZeroFill()

	return nil
}

func (f *Field) GetName() {
	ex := `\s{2}\x60(?P<Field>[0-9,a-z,A-Z$_\.]+)\x60`
	f.Name = common.FindMatchOne(ex, f.Raw, 1)
}

func (f *Field) GetUnsigned() {
	f.Unsigned = common.StringIn(f.Raw, " UNSIGNED")
}

func (f *Field) GetAutoIncrement() {
	f.AutoIncrement = common.StringIn(f.Raw, " AUTO_INCREMENT")
}

func (f *Field) GetZeroFill() {
	f.ZeroFill = common.StringIn(f.Raw, " ZEROFILL")
}

func (f *Field) GetNotNull() {
	f.NotNull = common.StringIn(f.Raw, " NOT NULL")
}

func (f *Field) GetType() {
	ex := `\s{2}\x60[0-9,a-z,A-Z$_\.]+\x60\s(?P<Type>\w+)`
	f.Type = strings.ToUpper(common.FindMatchOne(ex, f.Raw, 1))
}

func (f *Field) GetLength() {
	ex := `\s{2}\x60[0-9,a-z,A-Z$_\.]+\x60\s\w+\((?P<Length>\d+)\)`
	f.Length = common.StringToInt(common.FindMatchOne(ex, f.Raw, 1))
}

func (f *Field) GetEnum() {
	ex := `\s{2}\x60[0-9,a-z,A-Z$_\.]+\x60\s\w+(?:\((?P<List>'.+')\))`
	f.Enum = common.StringToArray(common.FindMatchOne(ex, f.Raw, 1))
}

func (f *Field) GetComment() {
	ex := `\sCOMMENT\s(?P<Comment>\w+|'(.*?)')[\s,]?`
	f.Comment = common.FindMatchOne(ex, f.Raw, 2)
}

func (f *Field) GetCollate() {
	ex := `\sCOLLATE\s(?P<Collate>\w+|'(.*?)')[\s,]?`
	f.Collate = common.FindMatchOne(ex, f.Raw, 2)
}

func (f *Field) GetDefault() {
	ex := `\sDEFAULT\s(?P<Default>\w+|'(.*?)')[\s,]?`
	f.Default = common.FindMatchOne(ex, f.Raw, 2)
}
