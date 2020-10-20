# go-fiber
Create an API with the same schema by various programming language, framework or library.

## Run
```bash
# start
$ go run main.go
```

## Structure
```
├── config.yml
├── go.mod
├── go.sum
├── main.go
└── src
    ├── auth
    │   ├── auth_controller.go
    │   └── auth_service.go
    ├── author
    │   ├── author_controller.go
    │   ├── author_repository.go
    │   └── author_service.go
    ├── database
    │   └── mysql.go
    ├── dto
    │   └── dto.go
    ├── middleware
    │   └── jwt.go
    └── user
        ├── user_controller.go
        ├── user_repository.go
        └── user_service.go
```
