# Build stage
FROM golang:1.23-alpine3.18 AS build

# Install necessary packages
RUN apk --no-cache add gcc g++ make ca-certificates git

# Set the working directory in the container
WORKDIR /go/src/github.com/Shridhar2104/logilo

# Copy go.mod and go.sum files and install dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the necessary application files
COPY . .

# Build the Go application
RUN GO111MODULE=on go build -o /go/bin/app ./shopify/cmd/shopify

# Runtime stage
FROM alpine:3.18

# Install necessary dependencies for running the Go app
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /usr/bin

# Copy the compiled binary from the build stage
COPY --from=build /go/bin/app .

# Expose the application port
EXPOSE 8080

# Command to run the Go application
CMD ["./app"]
