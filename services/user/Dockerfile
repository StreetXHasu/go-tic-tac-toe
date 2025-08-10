FROM golang:1-alpine AS builder

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY . .
RUN go mod download && go build -o user-service .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/user-service .
EXPOSE 8082

CMD ["./user-service"] 