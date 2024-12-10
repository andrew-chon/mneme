build:
	@echo "Building..."
	@cd server && go build -o mneme cmd/app/main.go

run:
	@cd server && go run cmd/app/main.go

clean:
	@echo "Cleaning..."
	@rm -f mneme