FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  

RUN apk --no-cache add \
    ca-certificates \
    curl \
    wget \
    jq

WORKDIR /root/

COPY --from=builder /app/main .
COPY avs.toml .
RUN chmod +x main

CMD ["./main","run"]
