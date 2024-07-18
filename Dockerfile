# Stage 1: Build Golang cli
FROM --platform=linux/amd64 golang:1.21-alpine AS go-build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o chainbase-cli .

# Stage 2: Setup Flink environment
FROM --platform=linux/amd64 flink:1.17

# Install necessary tools
USER root
RUN apt-get update && apt-get install -y curl

# Copy Golang binary
COPY --from=go-build /app/chainbase-cli /opt/chainbase-cli

# Copy entrypoint script
COPY entrypoint.sh /opt/entrypoint.sh
COPY avs.toml /opt/avs.toml
RUN chmod +x /opt/entrypoint.sh

# Switch back to the flink user
USER flink

ENTRYPOINT ["/opt/entrypoint.sh"]
