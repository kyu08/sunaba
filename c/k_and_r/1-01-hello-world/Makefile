ARG ?= main
TARGET = target

.PHONY: run
run:
	@clang $(ARG).c -o ./$(TARGET)/$(ARG) -std=c11 -Wall -Wextra \
	&& ./$(TARGET)/$(ARG)

.PHONY: lint
lint:
	cppcheck --quiet .
