//+build wireinject

package httpEngine

import (
	"github.com/google/wire"
	"github.com/lmx-Hexagram/gemini-generator/internal/service/config"
)

func InitDep() *service {
	wire.Build(
		config.ProvideHttp,
		wire.Struct(new(service), "*"),
	)

	//wire.Build(config.GetService().GetNekoMySQLConfig())
	return &service{}
}
