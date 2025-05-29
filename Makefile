.PHONY: build_and_start

build_and_start:
	go build -o lenslocked . && ./lenslocked
