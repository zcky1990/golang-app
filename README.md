# golang-app
This is Golang repository project,this repository used to practice and learning how to create application using GO

to run application: 
```
go run .
```

to run all unit test :
```
go test ./test/*
```

to run spesific unit test :
```
go test ./test/models
```

if you want to split environment variable to seperate file between production and test create new file with name `.env.test`

to run all unit test with .env.test:
```
ENV=test go test ./test/*
```

to run spesific unit test with .env.test:
```
ENV=test go test ./test/models
```