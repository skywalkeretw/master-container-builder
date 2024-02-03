# Start from golang base image
FROM golang:1.21 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod tidy

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

COPY openapi.json asyncapi.json ./

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]