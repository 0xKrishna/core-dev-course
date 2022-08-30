GO ?= latest
GOBIN = $(CURDIR)/build/bin
GORUN = env GO111MODULE=on go run
GOPATH = $(shell go env GOPATH)
GOTEST = GODEBUG=cgocheck=0 go test -p 1
TESTALL = $$(go list ./...)

lint:
	@./build/bin/golangci-lint run --config ./.golangci.yml

lintci-deps:
	rm -f ./build/bin/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./build/bin v1.46.0

test:
	$(GOTEST) --timeout 5m -shuffle=on -cover -coverprofile=cover.out $(TESTALL)
