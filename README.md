## Project Structure

```plaintext
go-fiber-boilerplate/
├── cmd/                        # main application
│   └── main.go
├── config/                     # load and manage config globally
│   └── config.go
├── internal/                   # internal packages
|   |── controllers/            # handle incoming HTTP request
|   |   └── user_controller.go
|   |── database/               
│   |   |── seeds/
|   |   └── migrations/
│   ├── middlewares/            # auth and etc
│   │   └── middleware.go
│   ├── models/                 # data models
│   │   └── model.go
│   |── services/               # handle business logic
│   |   └── user_service.go
├── pkg/                        # reusable external packages
│   ├── utils/
│   │   └── utils.go            # reusable function
│   └── logger/
│       └── logger.go           # logs
├── routes/                     # map incoming HTTP request
│       └── routes.go
├── .env
├── .gitignore
├── go.mod
├── go.sum
└── README.md