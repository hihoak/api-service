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

.PHONY: tests
tests:
	go test -v -race -parallel 8 -coverprofile out.txt ./...

.PHONY: generate-pb
generate-pb:
	protoc --proto_path=protos protos/hihoak/music_api/v1/*.proto --go_out=plugins=grpc:pkg/pb

.PHONY: remove-pb
remove-pb:
	rm -rf pkg/pb/*

.PHONY: super-linter
super-linter:
	docker run -e RUN_LOCAL=true -e DEFAULT_WORKSPACE=/Users/artemikhaylov/workspace/pet-project/api-service -v /Users/artemikhaylov/workspace/pet-project/api-service github/super-linter

help:
	@echo "hello - prints greetings with args GREETINGS_WORD, NAME and EXCLAMATION_POINT"
