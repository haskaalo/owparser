GO ?= go
PACKAGES ?= $(shell $(GO) list github.com/haskaalo/owparser/... | grep -v /vendor/)

lint:
	golint $(PACKAGES)

.PHONY: benchmark
benchmark:
	go test -bench . -benchmem