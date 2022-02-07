BOOX=boox

.PHONY: build build-dev run run-test

build:
	docker build -t boox .

build-dev:
	docker-compose build $(BOOX)

run:
	docker-compose up $(BOOX)

run-test:
	go test ./... -cover