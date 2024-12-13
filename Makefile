build:
	@echo "Building..."
	@go build -o mneme server/cmd/api/main.go

run:
	@go run server/cmd/api/main.go

clean:
	@echo "Cleaning..."
	@rm -f mneme

lint: 
	@golangci-lint run