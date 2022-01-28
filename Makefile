BINARY=wechat-work-pusher
VERSION=1.0.0
DATE=`date +%FT%T%z`
GoVersion=`go version`
LDFLAGS=-ldflags "-s -w -X main.version=${VERSION} -X 'main.date=${DATE}' -X 'main.goVersion=${GoVersion}'"
.PHONY: init build build_osx

default:
	@echo ${BINARY}
	@echo ${VERSION}
	@echo ${DATE}
	@echo ${GoVersion}

init:
	@go generate
	@echo "[ok] generate"

build:
	@GOOS=linux GOARCH=amd64 go build -o ${BINARY} ${LDFLAGS}
	@echo "[ok] build ${BINARY}"

build_arm:
	@GOOS=linux GOARCH=arm go build -o ${BINARY} ${LDFLAGS}
	@echo "[ok] build_arm ${BINARY}"

build_osx:
	@go build -trimpath -o ${BINARY} ${LDFLAGS}
	@echo "[ok] build_osx"

fmt:
	@gofmt -s -w ./