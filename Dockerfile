FROM golang:1.20-alpine

WORKDIR /app
COPY . .

RUN go build -o hello-app .

EXPOSE 8080

CMD ["./hello-app"]