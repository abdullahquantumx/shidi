# Build stage
FROM golang:1.23-alpine3.18 AS build

# Install necessary packages for building the Go application
RUN apk --no-cache add gcc g++ make ca-certificates git

# Set the working directory in the container
WORKDIR /go/src/github.com/Shridhar2104/logilo

# Copy go.mod and go.sum to install dependencies
COPY go.mod go.sum ./

# Install dependencies and clean up the Go modules
RUN go mod tidy

# Copy the necessary application files (e.g., 'account' folder) to the build container
COPY account account

# Build the Go application
RUN GO111MODULE=on go build -o /go/bin/app ./account/cmd/account

# Runtime stage
FROM alpine:3.18

# Install necessary runtime dependencies for the Go application
RUN apk --no-cache add ca-certificates

# Set the working directory in the runtime container
WORKDIR /usr/bin

# Copy the compiled binary from the build stage
COPY --from=build /go/bin/app .

# Expose the application port
EXPOSE 8080

# Command to run the Go application
CMD ["./app"]
