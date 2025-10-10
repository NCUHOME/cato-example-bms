package books

import (
	model "cato-example-bms/internal/models/database/books"
	repo "cato-example-bms/internal/repo/books"
	"context"
)

type engine struct{}

func (rdb *engine) FetchOne(ctx context.Context, table, sql string, args ...interface{}) (*model.Books, error) {
	//TODO implement me
	panic("implement me")
}

func (rdb *engine) FetchAll(ctx context.Context, table, sql string, args ...interface{}) ([]*model.Books, error) {
	panic("implement me")
}

func (rdb *engine) Exec(ctx context.Context, table, sql string, args ...interface{}) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (rdb *engine) Count(ctx context.Context, table, sql string, args ...interface{}) (int64, error) {
	//TODO implement me
	panic("implement me")
}

type RdbRepoBooks struct {
	engine
	basicRepo
}

func NewBooksRepo() repo.Repo {
	eg := engine{}
	basic := basicRepo{engine: eg}
	return &RdbRepoBooks{basicRepo: basic, engine: eg}
}
