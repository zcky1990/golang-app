# Golang App
This repository is a Go project designed for practicing and learning how to create a REST API application using Go. 
I'm follow the Ruby on Rails (ROR) folder structure as a base because I'm using ruby as main languange, I think it's fun to create this project using ruby structure folders

# JS and CSS Folder Structure
The JS folder structure is separated based on the environment you set:

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
see script section on package.json to know what command do  

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

see webpack.config.js to learn more  

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
# Development
When you develop using this, you need to run go and webpack at the same time,
it's better to use fresh, so you dont have to kill golang app every time you changes someting in your golang app but the choice is yours

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