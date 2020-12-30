package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// InstallmentEntity holds the schema definition for the InstallmentEntity entity.
type InstallmentEntity struct {
	ent.Schema
}

// Fields of the InstallmentEntity.
func (InstallmentEntity) Fields() []ent.Field {
	return []ent.Field{
		field.String(""),
		field.Time(""),
	}
}

// Edges of the InstallmentEntity.
func (InstallmentEntity) Edges() []ent.Edge {
	return nil
}
