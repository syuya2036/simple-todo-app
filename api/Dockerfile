# golang

FROM golang:1.21.1-alpine3.18

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./app/*

EXPOSE 8080

CMD ["./main"]