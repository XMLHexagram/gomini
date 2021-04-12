//+build wireinject

package db

import (
	"github.com/google/wire"
	"github.com/lmx-Hexagram/gemini-generator/internal/service/config"
)

func InitDep() *service {
	wire.Build(
		provideDbMap,
		config.ProvideDbMap,
		wire.Struct(new(service), "*"),
	)

	//wire.Build(config.GetService().GetNekoMySQLConfig())
	return &service{}
}
