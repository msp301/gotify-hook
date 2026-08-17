FROM golang:1.25 AS builder

WORKDIR /app
COPY . .
RUN make

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/gotify-hook .

RUN adduser -D -u 1000 appuser
USER appuser
EXPOSE 8080

ENTRYPOINT ["/app/gotify-hook"]
