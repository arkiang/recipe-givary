FROM golang:1.24-alpine

# Enable CGO for sqlite
ENV CGO_ENABLED=1
ENV GO111MODULE=on

# Install Alpine-compatible dependencies
RUN apk add --no-cache gcc musl-dev sqlite sqlite-dev

WORKDIR /app

# Copy go.mod and go.sum first (better caching for dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build from /main.go
RUN go build -o recipe main.go

FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache sqlite
COPY --from=builder /app/bin/app .
COPY --from=builder /app/migrations ./migrations
VOLUME /data
EXPOSE 3000
CMD ["./app"]