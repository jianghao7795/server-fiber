dev:
	go run .
start:
	air
build:
	go build main.go -c config.dev.yaml
mod:
	go mod tidy
main:
	go run main.go
run:
	go run .