# Golang App
This repository is a Go project designed for practicing and learning how to create a REST API application using Go. We follow the Ruby on Rails (ROR) folder structure as a base.

# Folder Structure
The folder structure is separated based on the environment you set:

Development: All JS and CSS files will be built in the static/dev folder.
Production: All JS and CSS files will be built in the static/dist folder.

# Build Tailwind CSS
To build Tailwind CSS, use the following commands:

Build Tailwind CSS:
````
npm run build:css
````
Build Tailwind CSS for Development/Test (not minified):
````
npm run build:css:dev
````
Build Tailwind CSS for Production (minified):
````
npm run build:css:prod
````
# Webpack
Start Webpack for hot reloading JS:

For Development:
````
npm run build:dev
````
For Production:
````
npm run build:prod
````
# Running the Go Application

Run the Go application with hot reload:
````
fresh
````

Run the Go application:
````
go run .
````
Run the Go application with a specific environment:
````
ENV=development go run .
ENV=production go run .
````

# Running Tests
Run all unit tests:
````
go test ./test/*
````
Run specific unit tests:
````
go test ./test/models
````
If you want to separate environment variables for production and test environments, create a new file named .env.test.

Run all unit tests with .env.test:
````
ENV=test go test ./test/*
````

Run specific unit tests with .env.test:
````
ENV=test go test ./test/models
````