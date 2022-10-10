GO=go
DOCKER=docker
BIN=main
SRC=main.go
PORT=8080
IMAGE=funfaker
CONTAINER=funfaker

.PHONY: build server server/localhost validate validate/autofix test test/verbose clean docker/build docker/run

build:
	$(GO) build -o $(BIN) $(SRC)

server: build
	./$(BIN) server -port $(PORT)

server/localhost: build
	./$(BIN) server -hostname localhost -port $(PORT)

validate: build
	./$(BIN) validate

validate/autofix: build
	./$(BIN) validate -autofix

test:
	$(GO) test ./...

test/verbose:
	$(GO) test -v ./...

clean:
	rm -f $(BIN)

docker/build:
	$(DOCKER) build --tag $(IMAGE) .

docker/run:
	$(DOCKER) run --name $(CONTAINER) --publish $(PORT):$(PORT)
