build:
	@echo "Building..."
	@go build -o mneme server/cmd/app/main.go

run:
	@go run server/cmd/app/main.go

clean:
	@echo "Cleaning..."
	@rm -f mneme