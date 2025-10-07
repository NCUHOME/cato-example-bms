package repo

import (
	"cato-example-bms/internal/models/database/books"
	"context"
	"database/sql"
)

type engine interface {
	FetchOne(ctx context.Context, sql string, args ...interface{}) (*books.Books, error)
	FetchAll(ctx context.Context, sql string, args ...interface{}) ([]*books.Books, error)
	Exec(ctx context.Context, sql string, args ...interface{}) sql.Result
}

type basic interface {
	engine

	FindById(ctx context.Context, id int64) (*books.Books, error)
	FindByName(ctx context.Context, name string) (*books.Books, error)
	FindByCategoryAndBookClass(ctx context.Context, category string, class string) ([]*books.Books, error)

	UpdateById(ctx context.Context, data map[string]interface{}, id int64) error
	UpdateByName(ctx context.Context, data map[string]interface{}, name string) (*books.Books, error)

	DeleteById(ctx context.Context, id int64) error
	DeleteByName(ctx context.Context, name string) error

	Insert(ctx context.Context, data *books.Books) error
}

type Repo interface {
	basic
	extra
}
