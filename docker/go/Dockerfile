# Stage 1 — build
FROM golang:1.26-alpine AS builder

WORKDIR /api

COPY api/go.mod api/go.sum ./
RUN go mod download

COPY api/ .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/server

# Stage 2 — run
FROM alpine:3.20

WORKDIR /app

COPY --from=builder /api/app ./app

EXPOSE 8080

CMD ["./app"]