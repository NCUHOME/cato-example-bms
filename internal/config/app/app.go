package app

import (
	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func init() {
	gin.SetMode(gin.DebugMode)
	app = gin.Default()
}

func GetApp() *gin.Engine {
	return app
}
