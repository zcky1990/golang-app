# golang-app
This is Golang repository project, this repository used to practice and learning how to create REST API Application using GO.
In this repository, we are following Ruby On Rails(ROR) stucture folder as base.

Build Tailwindcss
```
npm run build:css
```

To running application: 
```
go run .
```

To running all unit test :
```
go test ./test/*
```

To running spesific unit test :
```
go test ./test/models
```

If you want to split environment variable to seperate file between production and test create new file with name `.env.test`

To running all unit test with .env.test:
```
ENV=test go test ./test/*
```

To running spesific unit test with .env.test:
```
ENV=test go test ./test/models
```