package books

import (
	"context"
	"net/http"

	booksinspect "cato-example-bms/internal/inspect/books"
)

func NewBookManagerServiceTier() BookManageServiceTier {
	return new(bookManagerServiceTier)
}

type bookManagerServiceTier struct{}

func (b bookManagerServiceTier) BuildSearchBooksByCategoryV1Request(input *http.Request) (context.Context, *booksinspect.SearchBooksByCategoryRequest) {
	return input.Context(), &booksinspect.SearchBooksByCategoryRequest{}
}

func (b bookManagerServiceTier) WrapSearchBooksByCategoryV1Response(output http.ResponseWriter, resp *booksinspect.SearchBooksByCategoryResponse, bizErr error) {
	output.Header().Set("Content-Type", "application/json")
	output.Write([]byte("{\"books\": \"not found\"}"))
	output.WriteHeader(http.StatusOK)
}
