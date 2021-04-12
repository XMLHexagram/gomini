package httpEngine

import "github.com/gin-gonic/gin"

func Run() {
	err := httpEngineService.HttpEngine.Run(httpEngineService.Http.Port)
	if err != nil {
		panic(err)
	}
}

func Group(relativePath string, handlers ...gin.HandlerFunc) *gin.RouterGroup {
	return httpEngineService.HttpEngine.Group(relativePath, handlers...)
}
