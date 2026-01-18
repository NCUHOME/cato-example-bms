package books

import (
	"context"

	booksinspect "cato-example-bms/internal/inspect/books"
	domainbooks "cato-example-bms/internal/models/domain/books"
)

func NewBookManagerService() BookManageServiceService {
	return new(bookManagerService)
}

type bookManagerService struct {
}

func (b *bookManagerService) SearchBooksByCategoryV1(ctx context.Context, request *booksinspect.SearchBooksByCategoryRequest) (*booksinspect.SearchBooksByCategoryResponse, error) {
	// get repo
	//bookRepo := books.GetRepo()
	return &booksinspect.SearchBooksByCategoryResponse{
		Code:    "0",
		Message: "ok",
		Body: []*domainbooks.BookBrief{
			{
				Id:    0,
				Name:  request.Body.Category,
				Cover: "this is cover",
			},
		},
	}, nil
}
