package config

type Config struct {
	//Server `toml:"Server"`
	//DbMap  `toml:"DbMap"`
	//Log    `toml:"Log"`
	Base
	TLS
	Gemini
}

type Base struct {
	BaseUrl      string
	//LanguageCode string
	//Entry        string
	//AutoIndex    bool
}

type TLS struct {
	CertFile string
	KeyFile  string
}

type Gemini struct {
	DefaultLang     string
	AutoRedirect    bool
	AutoRedirectUrl string
	File            []struct {
		Router string `toml:"Router"`
		Path   string `toml:"Path"`
	} `toml:"File"`
	Dir []struct {
		Router        string `toml:"Router"`
		Path          string `toml:"Path"`
		Index         string `toml:"Index"`
		AutoCatalogue bool   `toml:"AutoCatalogue"`
	} `toml:"Dir"`
	Proxy []struct {
		Router    string `toml:"Router"`
		Method    string `toml:"Method"`
		URL       string `toml:"URL"`
		MetaField string `toml:"MetaField"`
		BodyField string `toml:"BodyField"`
	} `toml:"Proxy"`
}
//type Http struct {
//	Port    string `toml:"Port"`
//	Timeout int    `toml:"Timeout"`
//}
//
//type Server struct {
//	Http Http `toml:"Http"`
//}
//
//type Db struct {
//	Driver string `toml:"Driver"`
//	DSN    string `toml:"DSN"`
//}
//
//type DbMap map[string]Db
//
//type Log struct {
//	Level       string
//	ToStd       bool
//	LogRotate   bool
//	Development bool
//	Sampling    bool
//	Rotate      Rotate
//}
//
//type Rotate struct {
//	Filename   []string
//	MaxSize    int
//	MaxAge     int
//	MaxBackups int
//}
