package model

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        string    `bun:"user_id,pk" json:"id"`
	Name      string    `bun:"user_name,notnull,default:'anonymous'" json:"name"`
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"updated_at"`
}

func NewUser(name string) User {
	return User{
		ID:   GenerateIdentifier().Value(),
		Name: name,
	}
}
