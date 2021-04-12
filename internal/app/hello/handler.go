package hello

import (
	"github.com/gin-gonic/gin"
	"github.com/go-conflict/toolkit"
	"net/http"
)

func sayHello(c *gin.Context) {
	toolkit.SetJson(c, http.StatusOK, "hello go-conflict", nil)
}
