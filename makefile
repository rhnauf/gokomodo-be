run:
	go run ./cmd/app/

build:
	go build -C cmd/app -o ../../gokomodo-be.exe

build-run:
	go build -C cmd/app -o ../../gokomodo-be.exe && ./gokomodo-be.exe