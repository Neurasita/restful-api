GO=go

DIST_DIR=dist

.PHONY: build clean

build:
	@${GO} build -o ${DIST_DIR}/server cmd/server/main.go

clean:
	@rm -rf ${DIST_DIR}

test:
	@${GO} test ./...
