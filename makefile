tidy:
	go mod tidy && go mod vendor

run:
	go run ./cmd/alfred/main.go

hot-reload:
	air --build.cmd "go build -o bin/alfred cmd/alfred/main.go" --build.bin "./bin/alfred"

