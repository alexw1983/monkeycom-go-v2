build:
	@go tool templ generate
	@go build -o bin/monkecom-go-v2 cmd/web/main.go

test:
	@got test -v ./...

run: build
	@./bin/monkecom-go-v2

templ:
	@go tool templ generate