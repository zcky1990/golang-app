package environments

import (
	"fmt"
	"golang_app/golangApp/lib"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var _ EnvironmentConfiguration = (*DevelopmentEnvirontment)(nil)

type DevelopmentEnvirontment struct {
	Config *Config
}

func NewDevEnvConfiguration(env string) *DevelopmentEnvirontment {
	config := &Config{
		Env:                   env,
		EnvFilename:           ".env",
		EngineHtmlPath:        "./app/views",
		EnginePageType:        ".html",
		EngineViewsLayout:     "layouts/application",
		StaticAssetPath:       "./static",
		StaticAssetPublicPath: "/public",
	}
	return &DevelopmentEnvirontment{
		Config: config,
	}
}

func (c *DevelopmentEnvirontment) GetConfiguration() *Config {
	return c.Config
}

func (c *DevelopmentEnvirontment) LoadEvirontmentFile() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error get current directory", err)
		return
	}
	rootDir := lib.FindRootDir(currentDir)
	err = godotenv.Load(rootDir + "/" + c.Config.EnvFilename)
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
}

func (c *DevelopmentEnvirontment) GetJSFilePath() string {
	return fmt.Sprintf("%s/dev/javascript", c.Config.StaticAssetPublicPath)
}

func (c *DevelopmentEnvirontment) GetCSSFilePath() string {
	return fmt.Sprintf("%s/dev/css", c.Config.StaticAssetPublicPath)
}
