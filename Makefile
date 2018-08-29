deps: dep-tools
	@dep ensure -ve

dep-tools:
	go get -u github.com/ethereum/go-ethereum
	cd $(GOPATH)/src/github.com/ethereum/go-ethereum && go install ./cmd/clef

install:
	go install ./cmd/clefui/main.go

build:
	go build -gcflags "all=-N -l" -o ./build/clefui ./cmd/clefui/main.go

debug: build
	dlv --listen=:2345 --headless=true --api-version=2 exec ./build/clefui

.PHONY: build install deps