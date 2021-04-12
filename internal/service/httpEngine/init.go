package httpEngine

import (
	"github.com/lmx-Hexagram/gemini-generator/internal/service/config"
	"github.com/lmx-Hexagram/gemini-generator/internal/service/log"
	"github.com/gin-gonic/gin"
	"github.com/go-conflict/toolkit"
)

type service struct {
	HttpEngine *gin.Engine `wire:"-"`
	config.Http
}

var httpEngineService *service

func Init() {
	s := InitDep()

	s.HttpEngine = toolkit.DefaultHttpEngine()

	httpEngineService = s

	log.Info("httpEngineService init successfully")
}
