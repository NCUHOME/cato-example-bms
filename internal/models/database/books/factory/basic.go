package factory

import (
	"cato-example-bms/internal/models/database/books"
	"context"

	"github.com/Masterminds/squirrel"
	"xorm.io/builder"
)

type basicRepo struct {
	db
}

func (b *basicRepo) FindById(ctx context.Context, id int64) (*books.Books, error) {
	container := &books.Books{Id: id}
	cond := squirrel.Eq{container.GetIdCol(): id}
	sql, args, err := squirrel.Select(container.AllCols()).From(container.TableName()).Where(cond).ToSql()
	if err != nil {
		return nil, err
	}
	return b.FetchOne(ctx, sql, args)

}

func (b *basicRepo) FindByName(ctx context.Context, name string) (*books.Books, error) {
	container := &books.Books{Name: name}
	cond := builder.Eq{container.GetNameCol(): name}
	ok, err := b.db.Context(ctx).Table(container.TableName()).Where(cond).Get(container)
	if err != nil || !ok {
		return nil, err
	}
	return container, nil
}

func (b *basicRepo) FindByCategoryAndBookClass(ctx context.Context, category string, bookClass string) ([]*books.Books, error) {
	container := &books.Books{Category: category, BookClass: bookClass}
	cond := squirrel.And{
		// as key sort
		squirrel.Eq{container.GetCategoryCol(): category},
		squirrel.Eq{container.GetBookClassCol(): bookClass},
	}
	results := make([]*books.Books, 0)
	sql, args, err := squirrel.Select(container.AllCols()).From(container.TableName()).Where(cond).ToSql()
	if err != nil {
		return nil, err
	}
	return b.db.FetchAll(ctx, sql, args...)
}

func (b *basicRepo) UpdateById(ctx context.Context, data map[string]interface{}, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (b *basicRepo) UpdateByName(ctx context.Context, data map[string]interface{}, name string) (*books.Books, error) {
	//TODO implement me
	panic("implement me")
}

func (b *basicRepo) DeleteById(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}

func (b *basicRepo) DeleteByName(ctx context.Context, name string) error {
	//TODO implement me
	panic("implement me")
}

func (b *basicRepo) Insert(ctx context.Context, data *books.Books) error {
	//TODO implement me
	panic("implement me")
}
