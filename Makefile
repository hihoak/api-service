# TODO: think about how grep without closing dots, ubuntu grep ignoring lazy regexp
EXPECTED_GO_VERSION:=go1.17.
GO_VERSION_SHORT:=$(shell echo `go version` | grep -o -E "go([0-9]+\.){2}")
ifneq ($(EXPECTED_GO_VERSION),$(GO_VERSION_SHORT))
$(error Not valid go version. You have '$(GO_VERSION_SHORT)'. Expect '$(EXPECTED_GO_VERSION)')
endif

.PHONY: hello
hello:
	@echo "${GREETINGS_WORD}, ${NAME}${EXCLAMATION_POINT}"

.PHONY: build
build:
	go build -o bin/ -v cmd/main.go

help:
	@echo "hello - prints greetings with args GREETINGS_WORD, NAME and EXCLAMATION_POINT"
