.PHONY: build run test clean watch lint tidy

help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

build:
	@echo "Building..."
	@go build -o mneme server/cmd/api/main.go

run:
	@go run server/cmd/api/main.go

test:
	@echo "Testing..."
	@go test ./... -v

test/race:
	@echo "Testing..."
	@go test ./... -v -race

tidy:
	go mod tidy -v
	go fmt ./...

clean:
	@echo "Cleaning..."
	@rm -f mneme

lint:
	@golangci-lint run

watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi
