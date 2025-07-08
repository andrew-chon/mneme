
tidy:
	@echo "Tidying..."
	@rm -rf bin/

build: tidy
	@echo "Building..."
	@go build -o bin/api cmd/api/main.go
	
run:
	@go run cmd/api/main.go

test:
	@echo "Testing..."
	@go test ./... -v

lint:
	@echo "Linting..."
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


.PHONY: build run test lint