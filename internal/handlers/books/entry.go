package books

import (
	"github.com/gin-gonic/gin"

	"cato-example-bms/internal/config/app"
)

func init() {
	handler := NewBookManageServiceHttpHandler(
		NewBookManagerService(),
		NewBookManagerServiceTier(),
		NewBookMangerServiceContainer(),
	)
	srv := app.GetApp().Group(handler.GetRouterGroupBase())
	for path, runner := range handler.GetRoutersMap() {
		m, p, err := handler.DecodeKey(path)
		if err != nil {
			panic(err)
		}
		srv.Handle(m, p, gin.WrapF(runner))
	}
}
