BUILD_VERSION   := "v0.1"
BUILD_DATE      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD)


clean:
	rm -rf dist

install:
	go install

.PHONY : all docker buildx clean install

.EXPORT_ALL_VARIABLES:

GO111MODULE = on
GOPROXY = https://goproxy.cn
