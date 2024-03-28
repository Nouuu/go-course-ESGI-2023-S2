FROM golang:1.22-alpine AS BUILDER

WORKDIR /app

COPY . .

RUN go build -o app main.go

FROM alpine:latest

WORKDIR /app

COPY --from=BUILDER /app/app .

ENTRYPOINT ["/app/app"]