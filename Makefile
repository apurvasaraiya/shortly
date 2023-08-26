all: run

build:
	go build -o ./bin/shortly .

run: build
	./bin/shortly

.PHONY: all run
