package repository

import (
	"context"

	"github.com/mktbsh/golang-devcontainer/infrastructure"
	"github.com/mktbsh/golang-devcontainer/model"
	"github.com/uptrace/bun"
)

type UserRepository interface {
	FindAll() []model.User
	// FindByID(id string) model.User
	Create(user model.User) bool
	// Update(user model.User) (bool, error)
}

type userRepository struct {
	db *bun.DB
}

func NewUserRepository() UserRepository {
	return &userRepository{
		db: infrastructure.DB,
	}
}

func (r *userRepository) FindAll() []model.User {
	list := []model.User{}

	r.db.NewSelect().Model(&list).OrderExpr("u.created_at DESC").Scan(context.Background())

	return list
}

func (r *userRepository) Create(user model.User) bool {
	res, err := r.db.NewInsert().Model(&user).Exec(context.Background())
	if err != nil {
		return false
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false
	}

	return count > 0
}
