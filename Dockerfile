FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o helloGoApp .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/helloGoApp /usr/local/bin/helloGoApp

EXPOSE 5000

CMD ["helloGoApp"]
