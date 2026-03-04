.SILENT:

lint:
	golangci-lint run -v ./...

unit:
	go test -race -coverprofile cover.out $(shell go list ./...)
	go tool cover -html cover.out -o cover.html

TESTS = ./logcheck/testdata
test:
	$(foreach file, $(wildcard $(TESTS)/*/*), echo $(file); go run ./cmd/main.go $(file); echo -------------------;)

IMAGE_NAME = "test_docker"
docker-build:
	docker build -t $(IMAGE_NAME) .

docker-run:
	docker run -it --rm $(IMAGE_NAME)