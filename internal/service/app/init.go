package app

import (
	"github.com/lmx-Hexagram/gemini-generator/internal/app/hello"
	"github.com/lmx-Hexagram/gemini-generator/internal/service/log"
)

func Init() {
	hello.Init()

	log.Info("appService init successfully")
}
