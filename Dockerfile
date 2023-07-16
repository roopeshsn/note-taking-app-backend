FROM golang:1.20-alpine

WORKDIR /app

# Copy the source code to the container
COPY . .

RUN go get -d -v ./...

# Build the Go application
RUN go build -o main .

# Expose the port on which the Go application will run
EXPOSE 8080

# Start the Go application
CMD ["./main"]