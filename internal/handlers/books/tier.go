package books

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/ncuhome/cato/core/httpc"

	booksinspect "cato-example-bms/internal/inspect/books"
)

func NewBookMangerServiceContainer() httpc.Container {
	c := new(httpc.BasicMapContainer)
	c.Init()
	return c
}

func NewBookManagerServiceTier() BookManageServiceServiceTier {
	return new(bookManagerServiceTier)
}

type bookManagerServiceTier struct{}

func (b bookManagerServiceTier) BeforeSearchBooksByCategoryV1Request(input *http.Request) (context.Context, *booksinspect.SearchBooksByCategoryRequest) {
	ctx, cancel := context.WithTimeout(input.Context(), time.Second)
	defer cancel()
	return ctx, &booksinspect.SearchBooksByCategoryRequest{
		Body: &booksinspect.SearchBooksByCategoryRequestBody{Category: input.URL.Query().Get("category")},
	}
}

func (b bookManagerServiceTier) AfterSearchBooksByCategoryV1Response(output http.ResponseWriter, resp *booksinspect.SearchBooksByCategoryResponse, err error) {
	data, err := json.Marshal(resp)
	if err != nil {
		output.Write([]byte("Service error"))
		output.WriteHeader(503)
	} else {
		output.Write(data)
		output.WriteHeader(200)
	}
}
