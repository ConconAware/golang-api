#multi-stage builds
FROM golang:1.18.4 AS builder

WORKDIR /app
COPY go.mod .
COPY . .
RUN go build ./...

#RUN go test ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o app .


FROM alpine:latest AS production
COPY --from=builder /app .

CMD ["./app"]