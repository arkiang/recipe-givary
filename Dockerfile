ARG GO_VERSION=1.24
FROM golang:${GO_VERSION}-bookworm as builder

# Install dependencies needed for CGO and SQLite
# Use apt-get for Debian-based images
ENV CGO_ENABLED=1
RUN apt-get update && apt-get install -y gcc libc6-dev libsqlite3-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go mod tidy
# Compile the application. The output is /app/run-app
RUN go build -o run-app ./main.go

# Stage 2: Create the final, lightweight image
FROM debian:bookworm

# It's good practice to set a WORKDIR
WORKDIR /app

# --- THIS IS THE KEY FIX ---
# Copy the compiled binary from the 'builder' stage
COPY --from=builder /app/run-app .
COPY ./migrations ./migrations
RUN mkdir -p /app/data

# Expose the port the app runs on
EXPOSE 3000

# Set the command to run the executable
CMD ["./run-app"]