# Stage 1: Build the Go application
FROM golang:alpine as builder

WORKDIR /build
COPY . .

RUN go build -o main .

# Stage 2: Create a lightweight image with only the compiled binary
FROM alpine

RUN adduser -S -D -H -h /app appuser
USER appuser

WORKDIR /app
COPY --from=builder /build/main .

EXPOSE 3001

CMD ["./main"]
