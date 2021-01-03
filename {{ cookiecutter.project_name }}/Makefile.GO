# -*- mode: makefile -*-
GOPROXY={{cookiecutter.goproxy_uri}}

go-deps:
	@echo "Verifying and cleaning up dependencies for ${GOSERVICE}..."
	GOPROXY=$(GOPROXY) go mod tidy && go mod verify

go-build:
	@echo "Building '${GOSERVICE}'..."
	GOPROXY=$(GOPROXY) CGO_ENABLED=0 go build -o ${GOSERVICE} -v .

go-test:
	@echo "Running all ${GOSERVICE} tests.."
	SERVICE_ENV=test GOPROXY=${GOPROXY} go test -cover -v -race ./... -args -config-path=${CURDIR}/resources/config

go-run:
	SERVICE_ENV=development GOPROXY=${GOPROXY} SERVICE_PORT={{cookiecutter.service_port}} SERVICE_LOG=debug go run main.go server
