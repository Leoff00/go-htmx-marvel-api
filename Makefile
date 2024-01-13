TARGET_DIR=bin
GOBUILD=go build

BINARY_NAME=go_marvel_htmx

local:
	go run main.go

build:
	$(GOBUILD) -o $(TARGET_DIR)/$(BINARY_NAME) 

run:
	./$(TARGET_DIR)/$(BINARY_NAME)

clean:
	go clean
	rm -f $(TARGET_DIR)/$(BINARY_NAME)


release: build run

.PHONY: local build run release