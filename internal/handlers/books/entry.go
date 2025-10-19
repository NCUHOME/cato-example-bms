package books

import (
	"strings"

	"github.com/gin-gonic/gin"

	"cato-example-bms/internal/config/app"
)

func init() {
	handler := NewBookManageServiceHttpHandler(
		NewBookManagerService(),
		NewBookManagerServiceTier(),
		NewBookManageServiceContainer(),
	)
	srv := app.GetApp().Group(handler.GetRouterGroupBase())
	for path, runner := range handler.GetRoutersMap() {
		pattern := strings.SplitN(path, " ", 2)
		srv.Handle(pattern[0], pattern[1], gin.WrapF(runner))
	}
}
