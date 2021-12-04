.PHONY: hello
hello:
	@echo "${GREETINGS_WORD}, ${NAME}${EXCLAMATION_POINT}"

help:
	@echo "hello - prints greetings with args GREETINGS_WORD, NAME and EXCLAMATION_POINT"
