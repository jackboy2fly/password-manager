BINARY_NAME=pw-mgr

build:
	go build -o cmd/bin/${BINARY_NAME} cmd/main/main.go

run: build
	./cmd/bin/${BINARY_NAME}

clean:
	go clean
	rm -r cmd/bin
