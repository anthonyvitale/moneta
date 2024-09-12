BASEDIR ?= ${PWD}
WORKDIR ?= $(PWD)/.work
MOCKDIR ?= $(PWD)/mocks


.PHONY: all
all: auto test

.PHONY: tidy
tidy: auto
	go mod tidy

.PHONY: auto
auto:
	go generate -x ./...

.PHONY: test
test: auto
	go test ./...

.PHONY: cov
cov: auto
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

.PHONY: clean
clean:
	find "${BASEDIR}" -name "*.auto.go" -print | xargs rm -f
	go clean
	rm -f "${BASEDIR}/coverage.out"
	rm -rf "${MOCKDIR}"
