BINARY_NAME=helm-release
GOFILES=$(shell find . -name "*.go" -type f)
VERSION=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-ldflags "-X github.com/meithon/helm-release/cmd.Version=$(VERSION)"

.PHONY: all build clean install test lint release snapshot

all: build

build: $(GOFILES)
	go build $(LDFLAGS) -o $(BINARY_NAME)

clean:
	go clean
	rm -f $(BINARY_NAME)
	rm -rf dist/

install: build
	mv $(BINARY_NAME) /usr/local/bin/

test:
	go test ./...

lint:
	golangci-lint run

run: build
	./$(BINARY_NAME)

# Create a new release tag
release:
	@echo "Creating release $(VERSION)"
	git tag -a v$(VERSION) -m "Release v$(VERSION)"
	git push origin v$(VERSION)

# Create a snapshot release with GoReleaser
snapshot:
	goreleaser release --snapshot --clean
