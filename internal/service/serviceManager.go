package service

import (
	"github.com/lmx-Hexagram/gemini-generator/internal/service/app"
	"github.com/lmx-Hexagram/gemini-generator/internal/service/config"
	"github.com/lmx-Hexagram/gemini-generator/internal/service/db"
	"github.com/lmx-Hexagram/gemini-generator/internal/service/httpEngine"
	"github.com/lmx-Hexagram/gemini-generator/internal/service/log"
	"github.com/lmx-Hexagram/gemini-generator/internal/service/toolkit"
)

func Init() {
	config.Init()
	log.Init()
	toolkit.Init()
	db.Init()
	httpEngine.Init()
	app.Init()
}

func Run() {
	httpEngine.Run()
}
