package environments

type Config struct {
	Env                   string
	EnvFilename           string
	EngineHtmlPath        string
	EnginePageType        string
	EngineViewsLayout     string
	StaticAssetPath       string
	StaticAssetPublicPath string
	JavaScriptOutputPath  string
}

type EnvirontmentConfiguration interface {
	GetConfiguration() *Config
	LoadEvirontmentFile()
	GetJSFilePath() string
	GetCSSFilePath() string
}
