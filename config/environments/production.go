package environments

import (
	"fmt"
	"golang_app/golangApp/lib"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var _ EnvironmentConfiguration = (*ProductionEnvirontment)(nil)

type ProductionEnvirontment struct {
	Config *Config
}

func NewProdEnvConfiguration(env string) *ProductionEnvirontment {
	config := &Config{
		Env:                   env,
		EnvFilename:           ".env",
		EngineHtmlPath:        "./app/views",
		EnginePageType:        ".html",
		EngineViewsLayout:     "layouts/application",
		StaticAssetPath:       "./static",
		StaticAssetPublicPath: "/public",
	}
	return &ProductionEnvirontment{
		Config: config,
	}
}

func (c *ProductionEnvirontment) GetConfiguration() *Config {
	return c.Config
}

func (c *ProductionEnvirontment) LoadEvirontmentFile() {
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

func (c *ProductionEnvirontment) GetJSFilePath() string {
	return fmt.Sprintf("%s/dist/javascript", c.Config.StaticAssetPublicPath)
}

func (c *ProductionEnvirontment) GetCSSFilePath() string {
	return fmt.Sprintf("%s/dist/css", c.Config.StaticAssetPublicPath)
}
