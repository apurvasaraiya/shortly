all: run

build_race:
	go build -race -o ./bin/shortly .

build:
	go build -o ./bin/shortly .

run: build
	./bin/shortly

install_mockgen:
	go install github.com/golang/mock/mockgen@v1.6.0

dockerise:
	docker build -t shortly:latest .

mocks: install_mockgen
	mockgen -source=repository/repository.go -package mocks -destination=mocks/repository_mock.go
	mockgen -source=service/url_service.go -package mocks -destination=mocks/service_mock.go

test: mocks
	go test -count=1 ./...

.PHONY: all run
