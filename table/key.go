package table

type Key struct {
	Fields []string
	Name string
	Raw string `json:"-"`
}

func (k *Key) Parser(r string) error {
	(*k) = Key { Raw: r }

	k.GetName()
	k.GetFields()

	return nil
}

func (k *Key) GetName() {
	ex := `\s\x60(?P<Name>[0-9,a-z,A-Z$_\.]+)\x60\s\(.*\)`
	k.Name = findMatchOne(ex, k.Raw, 1)
}

func (k *Key) GetFields() {
	ex := `\s\x60[0-9,a-z,A-Z$_\.]+\x60\s\((?P<Fields>(\x60.+\x60(, )?)+)\)`
	k.Fields = stringToArray(findMatchOne(ex, k.Raw, 1))
}
