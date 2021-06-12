FROM golang:1.15.13-alpine3.13 AS base

ENV GO111MODULE=on

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build

FROM scratch
COPY --from=base /app/saber /app/

CMD ["/app/saber"]