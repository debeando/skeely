package index

import (
	"fmt"

	"mysql-ddl-lint/plugins/registry"
)

type Index struct {
	Property registry.Property
}

func init() {
	registry.Add("Index", func() registry.Method { return &Index{} })
}

func (f *Index) Run(p registry.Property) registry.Property {
	f.Property = p

	return f.Property
}

// - 701 Error: No duplicados.
// - Error: No confundir con unique key.
// - Error: Revisar indice correcto para texto.
// - Sugerir nombres valido.
// - Que el nombre termine en \_id
