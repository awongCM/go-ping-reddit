.PHONY: setup demo run build vet test clean

BINARY := bin/go-ping-reddit

setup:
	./scripts/setup.sh

build: $(BINARY)

$(BINARY): main.go demo.go validate.go go.mod go.sum
	mkdir -p bin
	go build -o $(BINARY) .

demo:
	go run . -demo

run:
	go run .

vet:
	go vet ./...

test:
	go test ./...

clean:
	rm -rf bin/
