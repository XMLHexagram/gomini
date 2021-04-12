package toolkit

import (
	"github.com/lmx-Hexagram/gemini-generator/internal/service/log"
	"github.com/go-conflict/toolkit"
)

func Init() {
	toolkit.Init(log.ProvideLogger())
}
