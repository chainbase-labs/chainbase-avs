FROM golang:1.22 AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod tidy && go mod verify

COPY . .

WORKDIR /usr/src/app/cli/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /usr/local/bin/cli .

FROM debian:latest

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=build /usr/local/bin/cli /usr/local/bin/cli

ENTRYPOINT [ "cli"]

CMD ["--config=/app/node.yaml"]