.SILENT:

lint:
	golangci-lint run -v ./...

unit:
	go test -race -coverprofile cover.out $(shell go list ./...)
	go tool cover -html cover.out -o cover.html

TESTS = ./logcheck/testdata/src
test:
	$(foreach file, $(wildcard $(TESTS)/*), echo $(file); go run ./cmd/main.go $(file); echo -------------------;)

plugin:
	golangci-lint custom -v

plugin-run: plugin
	./custom-gcl run -c ./.local.golangci.yml -v ./logcheck/testdata/src/slog ./logcheck/testdata/src/zap