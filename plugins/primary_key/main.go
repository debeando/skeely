package primary_key

import (
	"mysql-ddl-lint/plugins/registry"
)

type PrimaryKey struct {
	Property registry.Property
}

func init() {
	registry.Add("PrimaryKey", func() registry.Method { return &PrimaryKey{} })
}

func (pk *PrimaryKey) Run(p registry.Property) {
	pk.Property = p
	pk.Property.Code = 500
}

// - 501 No hay primary key.
// - 502 El nombre contiene puntos.
// - 503 No debe ser capital leter.
// - 504 El primary key debe ser not null.
// - 505 El primary key debe ser big int si es cualquiera numerico o char.
// - 506 El primary key debe ser char sino es bigint.
// - 507 El primary key debe ser autoincrement si el tipo de dato es int.
// - 508 El primary key no debe ser igual unique.
