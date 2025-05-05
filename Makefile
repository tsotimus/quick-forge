.PHONY: build run clean
.DEFAULT_GOAL := build

build:
	go build -o quickforge

run:
	./quickforge

clean:
	rm -f quickforge