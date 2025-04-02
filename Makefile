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

install-lint:
	$(GO) install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@2b224c2cf4c9f261c22a16af7f8ca6408467f338

lint: install-lint
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
