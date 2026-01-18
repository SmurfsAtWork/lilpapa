.PHONY: all build lilpapa

SERVER_BINARY_NAME=lilpapa

all: build-server

build: init build-server

build-server: init
	go build -ldflags="-w -s" -o ${SERVER_BINARY_NAME} ./cmd/server/...

init:
	go mod tidy

dev:
	@air -v && \
	air || \
	echo "air was not found, installing it..." && \
	go install github.com/cosmtrek/air@v1.51.0 &&  \
	air

lilpapa:
	./${MIGRATOR_BINARY_NAME} &&\
	./${SERVER_BINARY_NAME}

clean:
	go clean
	rm -f papa.db
