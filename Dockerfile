# syntax=docker/dockerfile:1

FROM golang:1.23.11-alpine AS builder

WORKDIR /go/src

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /app ./cmd/web

FROM golang:1.23.11-alpine AS runner

WORKDIR /

COPY --from=builder /app /app
COPY --from=builder /go/src/tls.sh /tls.sh

RUN chmod +x /tls.sh

ENTRYPOINT ["/tls.sh", "/app"]
CMD ["-addr=:", "-dsn="]
