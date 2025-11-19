
FROM golang:latest

RUN mkdir -p /app

WORKDIR /app
COPY go.mod /app
COPY go.sum /app
COPY .env /app
COPY . /app

RUN go mod download

RUN go build -o q-a ./cmd

EXPOSE 8080
CMD ["./q-a"]
