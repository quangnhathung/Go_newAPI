FROM golang:1.24

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o app

EXPOSE 3005

CMD ["./app"]
