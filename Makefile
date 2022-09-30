GO=go
BIN=main
SRC=main.go
PORT=8080
IMAGE=funfaker
CONTAINER=funfaker

build:
	$(GO) build -o $(BIN) $(SRC)

server: build
	./$(BIN) server -port $(PORT)

server/localhost: build
	./$(BIN) server -port $(PORT)

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
	docker build --tag $(IMAGE) .

docker/run:
	docker run --name $(CONTAINER) --publish $(PORT):$(PORT)
