run:
	go run ./cmd/app/

live-reload:
	nodemon --exec go run ./cmd/app/. --ext go
