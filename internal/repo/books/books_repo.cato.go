// code generate by proto-gen-cato, DO NOT EDIT

package books

import (
	model "cato-example-bms/internal/models/database/books"
	"context"
)

type basic interface {
	FindById(ctx context.Context, container *model.Books) (*model.Books, error)
	UpdateById(ctx context.Context, data map[string]interface{}, container *model.Books) error
	DeleteById(ctx context.Context, container *model.Books) error
	FindByName(ctx context.Context, container *model.Books) (*model.Books, error)
	UpdateByName(ctx context.Context, data map[string]interface{}, container *model.Books) error
	DeleteByName(ctx context.Context, container *model.Books) error
	FindByCategoryAndBookClass(ctx context.Context, container *model.Books) ([]*model.Books, error)
	Insert(ctx context.Context, data map[string]interface{}) error
}
