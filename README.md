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
```

## Running the Application

To run the application, you need to have Go installed on your machine.

```bash
# clone the repository
git clone https://github.com/senzel/go-fiber-boilerplate.git


# navigate to the project directory
cd go-fiber-boilerplate

# install dependencies
go mod download

# run the application
go run cmd/main.go
```

The application will start running on `http://localhost:3000`.

## Database

The application uses `GORM` as the ORM. The database is set up using `mySQL` by default.

To run locally, you need to pull mySQL docker image and run it.

```bash
# pull mySQL docker image
docker pull mysql:latest

# run mySQL docker container
docker run --name mysql -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=go-fiber-boilerplate -p 3306:3306 -d mysql:latest
```
