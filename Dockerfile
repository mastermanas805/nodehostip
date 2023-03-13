# Start with the official Go image
FROM golang:latest AS build

# Set the working directory
WORKDIR /app

# Copy the Go modules file
COPY go.mod .

# Download the dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mybinary .

# Start with a lightweight Alpine image
FROM alpine:latest

# Copy the binary from the previous build stage
COPY --from=build /app/mybinary .

# Run the binary
CMD ["./mybinary"]
