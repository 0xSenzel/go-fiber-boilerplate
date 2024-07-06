# Define the path variable

build:
	go build -o server cmd/main.go

# run: build
# 	clear
# 	./server

run: build
	clear
	go run cmd/main.go

watch:
	reflex -s -r '\.go$$' make run