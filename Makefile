BINARY=gotify-hook

.PHONY: build clean

build:
	CGO_ENABLED=0 go build -o $(BINARY)

clean:
	rm -f $(BINARY)
