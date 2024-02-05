clean-go/
├── cmd/
│   └── main.go
├── api/
│   └── api-spec.json (Postman API collection)
├── internal/
│   ├── delivery/
│   │   ├── http/
│   │   │   ├── middleware/
│   │   │   │   └── middleware.go (code for middleware)
│   │   │   ├── route/
│   │   │   │   └── route.go (initialize routes for server and run server)
│   │   │   └── user_controller.go (layer controller handling HTTP requests)
│   │   └── messaging/
│   ├── entity/
│   │   └── user_entity.go (declaration of user entity struct)
│   ├── usecase/
│   │   └── user_usecase.go (layer usecase handling business logic)
│   ├── repositories/
│   │   └── user_repository.go (layer repositories handling HTTP interactions to the database)
│   └── config/
│       └── db.go (initialize database connection)
|   └── gateway/
│       └── db.go (initialize database connection)
├── migration/
│   ├── User.go (initialize user table migration)
│   ├── migrate.go (setup migrate file)
│   └── token.go (initialize token table migration)
├── test/
│   └── user_test.go (unit testing here)
├── go.mod
├── makefile
├── .env
└── vendor/