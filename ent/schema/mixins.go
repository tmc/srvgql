package schema

import (
	"math/rand"
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect"

	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
	"github.com/oklog/ulid"
	"github.com/tmc/pulid"
)

// Mixin definitions

// AuditMixin implements the ent.Mixin for sharing
// audit-log capabilities with package schemas.
type AuditMixin struct {
	mixin.Schema
}

// Helpers for ID generation for entities.
type timeNowFn func() time.Time

type idGenerator struct {
	prefix string
	timeFn timeNowFn
}

func (idg *idGenerator) newULID() pulid.ID {
	timeFn := idg.timeFn
	if idg.timeFn == nil {
		timeFn = time.Now
	}
	t := timeFn()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := pulid.ID{Prefix: idg.prefix}
	id.ULID = ulid.MustNew(ulid.Timestamp(t), entropy)
	return id
}

func idField(prefix string) ent.Field {
	idGen := &idGenerator{prefix: prefix}
	return field.UUID("id", pulid.ID{}).
		Unique().
		SchemaType(map[string]string{
			dialect.MySQL:    "varchar(32)",
			dialect.Postgres: "text",
			dialect.SQLite:   "text",
		}).
		Default(idGen.newULID)
}

// Fields of the AuditMixin.
func (AuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Int("created_by").
			Optional(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Int("updated_by").
			Optional(),
	}
}

// TODO(tmc): Add hook to populate above fields.
