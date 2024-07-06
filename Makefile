# Define the path variable

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