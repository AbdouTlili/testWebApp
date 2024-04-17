# Use the official Golang image as a base image
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

# Build the  app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o main .

FROM scratch

COPY --from=builder /app/main /app/main

EXPOSE 8080

# Command to run the executable
CMD ["/app/main"]
