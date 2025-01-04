FROM golang:1.23.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go mod and sum files for dependency installation
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN make build-prod

# Step 2: Use a smaller base image for the final container
FROM debian:bookworm-slim AS prod

# Set a working directory
WORKDIR /app

# Copy the built Go binary from the builder stage
COPY --from=builder /app/mneme .

# Expose the port the application will listen on
EXPOSE 8080

# Command to run the executable
CMD ["./mneme"]
