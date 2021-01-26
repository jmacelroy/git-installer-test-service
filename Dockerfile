FROM golang:buster as builder

WORKDIR /app
ADD . .
RUN go build -o /usr/local/bin/test-service

EXPOSE 8080
CMD ["/usr/local/bin/test-service"]
