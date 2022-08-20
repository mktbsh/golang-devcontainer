package model

import (
	"crypto/rand"
	"io"
	"time"

	"github.com/oklog/ulid/v2"
)

type Identifier struct {
	id string
}

type IdentifierGenerater interface {
	Generate() Identifier
}

var defaultGenerator IdentifierGenerater

func init() {
	defaultGenerator = newULIDGenerator(rand.Reader)
}

func GenerateIdentifier() Identifier {
	return defaultGenerator.Generate()
}

func NewIdentifier(id string) Identifier {
	return Identifier{id}
}

func (i Identifier) Value() string {
	return i.id
}

func (i Identifier) Equal(other Identifier) bool {
	return i.id == other.id
}

type ULIDGenerator struct {
	entropy *ulid.MonotonicEntropy
}

func newULIDGenerator(r io.Reader) *ULIDGenerator {
	return &ULIDGenerator{
		entropy: ulid.Monotonic(r, 0),
	}
}

func (g *ULIDGenerator) Generate() Identifier {
	id := ulid.MustNew(ulid.Timestamp(time.Now()), g.entropy)
	return Identifier{id: id.String()}
}
