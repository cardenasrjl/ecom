build:
	go build -o server cmd/server/main.go

build_client:
	go build -o client cmd/client-grpc/main.go

up: build
	docker-compose up -d
	./server
