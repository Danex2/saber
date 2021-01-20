FROM golang:1.15.7-alpine3.13

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go mod download

RUN go build -o main .

CMD ["/app/main"]

