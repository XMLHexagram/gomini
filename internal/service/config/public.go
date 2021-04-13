package config

func ProvideTLS() TLS {
	return configService.TLS
}

func ProvideBase() Base {
	return configService.Base
}

func ProvideGemini() Gemini {
	return configService.Gemini
}
