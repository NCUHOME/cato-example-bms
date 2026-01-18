package engine

import (
	"context"
)

type MockEngine[T any] struct {
}

func (m *MockEngine[T]) FetchOne(ctx context.Context, table string, sql string, args ...interface{}) (*T, error) {
	return new(T), nil
}

func (m *MockEngine[T]) FetchAll(ctx context.Context, table interface{}, sql string, args ...interface{}) ([]*T, error) {
	return make([]*T, 0), nil
}

func (m *MockEngine[T]) Exec(ctx context.Context, table, sql string, args ...interface{}) (int64, error) {
	return 0, nil
}

func (m *MockEngine[T]) Count(ctx context.Context, table, sql string, args ...interface{}) (int64, error) {
	return 0, nil
}
