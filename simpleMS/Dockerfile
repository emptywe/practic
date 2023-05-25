FROM golang:1.20-alpine AS builder

COPY . /simpleMS/
WORKDIR /simpleMS/

RUN go mod download

RUN go build -o ./bin/simpleMS main.go

FROM alpine:3.17
RUN apk --update --no-cache add ca-certificates tzdata
WORKDIR /usr/src/app
COPY --from=builder /simpleMS/bin/simpleMS .

ENTRYPOINT ["./simpleMS"]

EXPOSE 5300