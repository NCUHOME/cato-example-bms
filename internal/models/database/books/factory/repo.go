package factory

import (
	"cato-example-bms/internal/models/database/books"
	"context"
	"database/sql"
)

type db struct{}

func (db *db) FetchOne(ctx context.Context, sql string, args ...interface{}) (*books.Books, error) {
	//TODO implement me
	panic("implement me")
}

func (db *db) FetchAll(ctx context.Context, sql string, args ...interface{}) ([]*books.Books, error) {
	panic("implement me")
}

func (db *db) Exec(ctx context.Context, sql string, args ...interface{}) sql.Result {
	//TODO implement me
	panic("implement me")
}

type BooksRepo struct {
	basicRepo
}

func (e BooksRepo) CalculateCount(ctx context.Context, id int64) (int64, error) {
	//TODO implement me
	panic("implement me")
}
