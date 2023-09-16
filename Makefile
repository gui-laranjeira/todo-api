build:
	go build -o bin/main main.go

run:
	go run main.go

compíle:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin main.go
