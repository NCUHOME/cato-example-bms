package books

import (
	"net/http"
)

func NewBookManageServiceContainer() BookManageServiceServiceContainer {
	return &routerContainer{mmap: make(map[string]http.HandlerFunc)}
}

type routerContainer struct {
	mmap map[string]http.HandlerFunc
}

func (r routerContainer) Set(method string, path string, runner func(w http.ResponseWriter, r *http.Request)) {
	r.mmap[method+path] = runner
}

func (r routerContainer) Get(method string, path string) (func(w http.ResponseWriter, r *http.Request), bool) {
	runner, ok := r.mmap[method+path]
	return runner, ok
}

func (r routerContainer) ToMap() map[string]http.HandlerFunc {
	return r.mmap
}
