package table

import (
	"regexp"
	"strings"
)

func (t *Table) RawString() string {
	return string(t.Raw)
}

func (t *Table) Parser(r []byte) bool {
	(*t) = Table { Raw: r }

	t.parserTable()
	t.parserFields()
	t.parserPrimaryKey()
	t.parserUniqueKeys()
	t.parserIndexes()
	t.parserConstraints()
	
	if len(t.Name) == 0 {
		return false
	}

	return true
}

func (t *Table) parserTable() {
	ex := `CREATE TABLE \x60(?P<Table>[0-9,a-z,A-Z$_\.]+)\x60 (?P<Body>\([^*]*\))(\s+)?(?:ENGINE=(?P<Engine>\w+))?(\s+)?(?:DEFAULT CHARSET=(?P<Charset>\w+))?(\s+)?(?:COLLATE=(?P<Collate>\w+))?(\s+)?(?:ROW_FORMAT=(?P<RowFormat>\w+))?(\s+)?(?:COMMENT='(?P<Comment>.+)')?`
	re := regexp.MustCompile(ex)
	match := re.FindStringSubmatch(t.RawString())

	for i, name := range re.SubexpNames() {
		if len(match) > 0 && i != 0 && name != "" {
			switch name {
				case "Table":
					t.Name = match[i]
				case "Engine":
					t.Engine = match[i]
				case "Charset":
					t.Charset = match[i]
				case "Collate":
					t.Collate = match[i]
				case "RowFormat":
					t.RowFormat = match[i]
				case "Comment":
					t.Comment = match[i]
			}
		}
	}
}

func (t *Table) parserFields() {
	for _, item := range find(`\s{2}\x60(?P<Name>[0-9,a-z,A-Z$_\.]+)\x60\s(?P<Properties>.*)`, t.RawString()) {
		t.Fields = append(t.Fields, Field{
			Name: item[1],
			// Parse properties, next implementation...
		})
	}
}

func (t *Table) parserPrimaryKey() {
	for _, item := range find(`\s{2}PRIMARY KEY \((?P<Fields>(\x60.+\x60(, )?)+)\)`, t.RawString()) {
		t.PrimaryKey = parserValuesAsStringArray(item[2])
		return
	}
}

func (t *Table) parserUniqueKeys() {
	for _, item := range find(`\s{2}UNIQUE KEY \x60(?P<Name>[0-9,a-z,A-Z$_\.]+)\x60 \((?P<Fields>(\x60.+\x60(, )?)+)\)`, t.RawString()) {
		t.UniqueKeys = append(t.UniqueKeys, Key{
			Name: item[1],
			Fields: parserValuesAsStringArray(item[2]),
		})
	}
}

func (t *Table) parserIndexes() {
	for _, item := range find(`\s{2}KEY \x60(?P<Name>[0-9,a-z,A-Z$_\.]+)\x60 \((?P<Fields>(\x60.+\x60(, )?)+)\)`, t.RawString()) {
		t.Keys = append(t.Keys, Key{
			Name: item[1],
			Fields: parserValuesAsStringArray(item[2]),
		})
	}
}

func (t *Table) parserConstraints() {
	for _, item := range find(`\s{2}CONSTRAINT \x60(?P<Name>[0-9,a-z,A-Z$_\.]+)\x60\s(?P<Properties>.*)`, t.RawString()) {
		t.Constraints = append(t.Constraints, Constraint{
			Name: item[1],
			// Parse properties, next implementation...
		})
	}
}

func parserValuesAsStringArray(v string) (values []string) {
	for _, value := range strings.Split(v, ",") {
		values = append(values, strings.Trim(value, "`"))
	}

	return values
}

func find(e string, t string) [][]string {
	return regexp.MustCompile(e).FindAllStringSubmatch(t, -1)
}
