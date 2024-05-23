package environments

// Config variable for Fiber template
type Config struct {
	// Environment variable used by the app
	// Values for this field are "development", "production", and "test"
	// You can create a staging environment if you want
	Env string

	// Filename of the .env file, like .env or .env.test
	// We are using dotenv to load env variables
	EnvFilename string

	// Path of the template file that we use in the application
	// Such as layout and view; in this application, we are using ./app/views as the default
	EngineHtmlPath string

	// Extension of the template that we use as layout and view
	// In this application, we are using .html as the extension of our template
	EnginePageType string

	// Default template layout we use
	// In this application, we are using layout/application inside views/layout as the default
	EngineViewsLayout string

	// Path of current assets, like JS, images, or CSS, that need to be mounted
	// In this application, we are using the ./static folder as our base static path
	StaticAssetPath string

	// Path of the static folder we can access via HTTP in Go; usually, we copy or save our JS, images, and CSS to this folder
	// If we don't use static, we can't access it via the browser
	// In this application, we are using /public as base public path
	// So we can access our JS and CSS using http://hostname/public/{javascript/css/anything}.*
	StaticAssetPublicPath string
}

// We make development, test, and production implement this interface
// So every environment will have different configurations and output paths for JS and CSS that we can import
type EnvironmentConfiguration interface {
	GetConfiguration() *Config
	LoadEnvironmentFile()
	GetJSFilePath() string
	GetCSSFilePath() string
}
