# Define the path variable
GRAPHQL_PATH := internal/database/graphql

build:
	go build -o server main.go

# run: build
# 	clear
# 	./server

run:
	clear
	go run cmd/main.go

watch:
	reflex -s -r '\.go$$' make run

gqlgen-generate:
	cd $(GRAPHQL_PATH) && go run github.com/99designs/gqlgen generate

run-graphql:
	cd $(GRAPHQL_PATH) && go run server.go