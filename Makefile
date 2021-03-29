build:
	mkdir -p deploy
	glide get github.com/labstack/echo
	go run main.go
