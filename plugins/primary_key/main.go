package primary_key

import (
	"fmt"

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

	pk.Empty()

	for _, message := range pk.Property.Messages {
		fmt.Println(fmt.Sprintf("- [%d] %s", pk.Property.Code + message.Code, message.Message))
	}
}

func (pk *PrimaryKey) AddMessage(id int, m string) {
	pk.Property.Messages = append(pk.Property.Messages, registry.Message{Code: id, Message: m})
}

func (pk *PrimaryKey) Empty() {
	if len(pk.Property.Table.PrimaryKey) == 0 {
		pk.AddMessage(1, "Table no have Primary Key.")
	}
}

// - 504 El primary key debe ser not null.
// - 505 El primary key debe ser big int si es cualquiera numerico o char.
// - 506 El primary key debe ser char sino es bigint.
// - 507 El primary key debe ser autoincrement si el tipo de dato es int.
// - 508 El primary key no debe ser igual unique.
