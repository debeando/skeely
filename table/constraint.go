package table

type Constraint struct {
	Delete string
	Fields []string
	Name string
	Raw string `json:"-"`
	Reference []Key
	Update string
}
