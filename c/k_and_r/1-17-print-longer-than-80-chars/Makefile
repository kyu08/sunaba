ARG ?= main
TARGET = target

.PHONY: run
run:
	-@mkdir -p $(TARGET) \
	&& clang $(ARG).c -o ./$(TARGET)/$(ARG) -std=c23 -Wall -Wextra \
	&& ./$(TARGET)/$(ARG)

.PHONY: lint
lint:
	cppcheck --quiet .
