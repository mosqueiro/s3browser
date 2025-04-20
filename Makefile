.PHONY: build
build:
	go build -o bin/s3browser

.PHONY: run
run:
	go run

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -race -cover ./...

.PHONY: build-image
build-image:
	docker build -t s3browser .

.PHONY: clean
clean:
	rm -rf bin
