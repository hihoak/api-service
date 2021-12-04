EXPECTED_GO_VERSION:=go1.17
GO_VERSION_SHORT:=$(shell echo `go version` | egrep -o "go(([0-9]+\.*?){2})")
ifneq ($(EXPECTED_GO_VERSION), $(GO_VERSION_SHORT))
$(error Not valid go version. You have '$(GO_VERSION_SHORT)'. Expect '$(EXPECTED_GO_VERSION)')
endif

.PHONY: hello
hello:
	@echo "${GREETINGS_WORD}, ${NAME}${EXCLAMATION_POINT}"

.PHONY: build
build:
	export SOURCE_PATH="${SOURCE_FOLDER}/${SOURCE_FILE_NAME}" \
	export OUTPUT_PATH="${DST_PATH}/${OUTPUT_FILE_NAME}" \
	go build -v -o ${OUTPUT_PATH} ${SOURCE_PATH}

help:
	@echo "hello - prints greetings with args GREETINGS_WORD, NAME and EXCLAMATION_POINT"
