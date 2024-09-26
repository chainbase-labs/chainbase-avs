FROM golang:1.22 AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod tidy && go mod verify

COPY . .

WORKDIR /usr/src/app/coordinator/cmd
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /usr/local/bin/coordinator .

FROM debian:latest
COPY --from=build /usr/local/bin/coordinator /usr/local/bin/coordinator
ENTRYPOINT [ "coordinator"]
CMD ["--config=/app/coordinator.yaml"]