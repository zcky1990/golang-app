package environments

// config variable for fibers template
type Config struct {
	// Environment variable used by app
	// value of this fields is developement, production and test
	// you can create staging env if you want
	Env string

	// Filename of .env, like .env or .env.test
	// we are using dotenv to load env
	EnvFilename string

	// path of template file that we use in application
	// such as layout and view, in this application we are using ./app/views as default
	EngineHtmlPath string

	//extension of template that we uese as layout and view
	//in this application we are using .html as extension of our template
	EnginePageType string

	// default template layout we use
	// in this application we are using layout/application inside views/layout as default
	EngineViewsLayout string

	// path of current asset, like js, image or css, we need to mounted
	// in this application we are using .static folder as our base static path
	StaticAssetPath string

	// path of static folder we can access via http in GO, usually we copy or save our js,image,css to this folder
	// if we dont use static, we cant access it via browser,
	// in this application we are using /public
	// so we can access out js and css using http://hostname/public/{javascript/css/anything}.*
	StaticAssetPublicPath string
}

// we make development,test and production implement's this interface
// so every environment will have diffrent configuration and outpout path of js and css that we can import
type EnvironmentConfiguration interface {
	GetConfiguration() *Config
	LoadEvirontmentFile()
	GetJSFilePath() string
	GetCSSFilePath() string
}
