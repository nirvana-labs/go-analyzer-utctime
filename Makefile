GO := go
GOFLAGS := -v
COVERPROFILE := coverage.out

clean:
	$(GO) clean -testcache
	rm -f $(COVERPROFILE)

tidy:
	$(GO) mod tidy

build: tidy
	$(GO) build $(GOFLAGS) ./...

vet:
	$(GO) vet $(GOFLAGS) ./...

fmt:
	$(GO) fmt ./...

lint:
	golangci-lint run --config .golangci.yaml

test:
	$(GO) test $(GOFLAGS) -race -coverprofile=$(COVERPROFILE) ./...
	$(GO) tool cover -func=$(COVERPROFILE)

install-license-check:
	$(GO) install github.com/google/go-licenses@5348b744d0983d85713295ea08a20cca1654a45e

license-check: install-license-check
	go-licenses check ./... --disallowed_types=forbidden,restricted --ignore=github.com/nirvana-labs

install-security-check:
	$(GO) install github.com/securego/gosec/v2/cmd/gosec@d4617f51baf75f4f809066386a4f9d27b3ac3e46

security-check: install-security-check
	gosec ./...

.PHONY: clean tidy build vet fmt lint test license-check security-check
