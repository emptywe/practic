PROJECTNAME=$(shell basename "$(PWD)")

.PHONY:
.SILENT:

build:
	go build -o cli cmd/cli/cli.go
	go build -o srv cmd/http.Srv/srv.go
run:
	go build -o cli cmd/cli/cli.go
	go run cmd/http.Srv/srv.go
docker:
	docker build -t clisub .
install:
	go install cmd/cli/cli.go
	go run cmd/http.Srv/srv.go