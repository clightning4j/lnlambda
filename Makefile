CC=go
FMT=gofmt
NAME=lambda
BASE_DIR=/script
OS=linux
ARCH=386
ARM=
GORPC_COMMIT=52b0b2cd43735132e59da994177f4242d51cae1a

default: fmt
	$(CC) build -o $(NAME) main.go

fmt:
	$(CC) fmt ./...

lint:
	golangci-lint run

check:
	$(CC) test -v ./...

check-dev:
	richgo test ./... -v

build:
	env GOOS=$(OS) GOARCH=$(ARCH) GOARM=$(ARM) $(CC) build -o $(NAME)-$(OS)-$(ARCH) main.go

run:
	$(CC) run main.go

dep:
	$(CC) mod vendor
