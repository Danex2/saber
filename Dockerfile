FROM golang:1.15.13-alpine3.13 AS base

WORKDIR /app/
COPY . .
RUN go mod download
RUN go build -o /build

FROM golang:1.15.13-alpine3.13
COPY --from=base /build /build


CMD ["./build/saber"]