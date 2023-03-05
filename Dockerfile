FROM golang:1.20.1-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN ls
RUN go build -o /app/build/go-api-fiber
FROM alpine:latest AS production
COPY --from=builder /app/build .
COPY --from=builder /app/.env .
ENTRYPOINT ["./go-api-fiber"]