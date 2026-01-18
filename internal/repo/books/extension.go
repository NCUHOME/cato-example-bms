package books

import (
	model "cato-example-bms/internal/models/database/books"
	"cato-example-bms/internal/rdb/books"
	"cato-example-bms/internal/repo/engine"
)

func init() {
	core := new(engine.MockEngine[model.Books])
	repo = books.NewRdbRepoBooks(core)
}

type extension interface{}
