run:
	@go run main.go

build:
	@go build -o ./dist/health main.go

build-rpi:
	@env GOOS=linux GOARCH=arm GOARM=7 go build -o ./dist/health main.go

