build:
	@cd cmd
	@go mod tidy
	@go build -o bin/password_gen cmd/main.go

run: build
	@./bin/password_gen
