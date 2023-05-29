FROM golang:1.20-alpine AS builder

COPY . /substirng/
WORKDIR /substirng/

RUN go build -o ./bin/srv ./cmd/httpSrv/srv.go
RUN go install ./cmd/cli/cli.go

CMD ["./bin/srv"]
