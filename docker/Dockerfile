# Stage 1: Build the Go binary
FROM golang:1.24.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN go build -o main ./

#RUN apk --no-cache add ca-certificates
# Stage 2: Create a lightweight runtime image 
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]