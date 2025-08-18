FROM golang:trixie AS builder

COPY . .
RUN go build -o wplace-bot ./cmd/bot

FROM debian:trixie

RUN mkdir /app
WORKDIR /app
COPY --from builder ./wplace-bot /app/wplace-bot
COPY config.json /app/config.json
ENTRYPOINT ["/app/wplace-bot"]