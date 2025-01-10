FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Set environment variables for the build
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

# Build the Go application
RUN go build -o app ./randomizer/cmd

# Copy .env file to build stage
COPY .env .env

# Final stage
FROM alpine:latest

WORKDIR /root/

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/app .

# Ensure the binary is executable
RUN chmod +x ./app

# Copy the .env file to the final image
COPY --from=builder /app/.env .env

# Copy the ABI files to the final image
COPY common/blockchain/abis common/blockchain/abis

CMD ["./app"]