package environments

import (
	"fmt"
	"golang_app/golangApp/lib"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var _ EnvironmentConfiguration = (*TestEnvirontment)(nil)

type TestEnvirontment struct {
	Config *Config
}

func NewTestEnvConfiguration(env string) *TestEnvirontment {
	config := &Config{
		Env:                   env,
		EnvFilename:           ".env",
		EngineHtmlPath:        "./app/views",
		EnginePageType:        ".html",
		EngineViewsLayout:     "layouts/application",
		StaticAssetPath:       "./static",
		StaticAssetPublicPath: "/public",
	}
	return &TestEnvirontment{
		Config: config,
	}
}

func (c *TestEnvirontment) GetConfiguration() *Config {
	return c.Config
}

func (c *TestEnvirontment) LoadEnvironmentFile() {
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

func (c *TestEnvirontment) GetJSFilePath() string {
	return fmt.Sprintf("%s/dev/javascript", c.Config.StaticAssetPublicPath)
}

func (c *TestEnvirontment) GetCSSFilePath() string {
	return fmt.Sprintf("%s/dev/css", c.Config.StaticAssetPublicPath)
}
