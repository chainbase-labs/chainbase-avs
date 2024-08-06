FROM alpine:latest

RUN apk --no-cache add \
    ca-certificates \
    curl \
    wget \
    jq

WORKDIR /root/

COPY ./main_amd64 .
COPY ./main_arm64 .
COPY ./entrypoint.sh .

RUN chmod +x entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
