
BINARY="idx"
OUTPUT_DIR=.
GOOS=$(shell go env GOOS)
#GOOS
APP_NAME="Identity.X"

BUILD_TIME=$(shell date "+%FT%T%z")
GIT_REVISION=$(shell git rev-parse --short HEAD)
GIT_BRANCH=$(shell git name-rev --name-only HEAD)
APP_VERSION=0.0.2
BUILD_VERSION=$(APP_VERSION).$(GIT_REVISION)
GO_VERSION=$(shell go version)
release:	
	cd cmd &&CGO_ENABLED=0 go build -a -v \
	-ldflags "-extldflags '-static' -s -X 'main.AppName=${APP_NAME}' \
				-X 'main.AppVersion=${APP_VERSION}' \
				-X 'main.BuildVersion=${BUILD_VERSION}' \
				-X 'main.BuildTime=${BUILD_TIME}' \
				-X 'main.GitRevision=${GIT_REVISION}' \
				-X 'main.GitBranch=${GIT_BRANCH}' \
				-X 'main.GoVersion=${GO_VERSION}' " -o ${OUTPUT_DIR}/${BINARY} .
	./cmd/${BINARY} -ver
docker:
	docker build . -t idx:dev --build-arg APP_VERSION=${APP_VERSION} 