package hello

import (
	"github.com/lmx-Hexagram/gemini-generator/internal/service/httpEngine"
)

func Init() {
	hello := httpEngine.Group("/hello")
	hello.GET("", sayHello)
}
