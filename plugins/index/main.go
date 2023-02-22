package index

import (
	"fmt"

	"mylinter/plugins/registry"
)

type Plugin struct {
	Property registry.Property
}

func init() {
	registry.Add("Index", func() registry.Method { return &Plugin{} })
}

func (p *Plugin) Run(p registry.Property) registry.Property {
	p.Property = p

	return p.Property
}

// - 701 Error: No duplicados.
// - Error: No confundir con unique key.
// - Error: Revisar indice correcto para texto.
// - Sugerir nombres valido.
// - Que el nombre termine en \_id
