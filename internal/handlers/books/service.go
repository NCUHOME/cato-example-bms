package books

import (
	"context"

	booksinspect "cato-example-bms/internal/inspect/books"
)

func NewBookManagerService() BookManageServiceService {
	return new(bookManagerService)
}

type bookManagerService struct {
}

func (b *bookManagerService) SearchBooksByCategoryV1(ctx context.Context, request *booksinspect.SearchBooksByCategoryRequest) (*booksinspect.SearchBooksByCategoryResponse, error) {
	return &booksinspect.SearchBooksByCategoryResponse{}, nil
}
