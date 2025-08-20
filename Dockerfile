FROM golang:tip-alpine3.22 AS builder

COPY . .

RUN go build -ldflags="-s -w" -o /tmp/ytgy

FROM alpine:3.22

COPY --from=builder /tmp/ytgy /usr/bin/ytgy

EXPOSE 8080

CMD ["/usr/bin/ytgy"]
