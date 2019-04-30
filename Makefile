setup: dep-tools
	qtsetup check
	qtsetup

deps: dep-tools
	@dep ensure -v
	@echo "Un-vendoring QT bindings..."
	@rm -rf vendor/github.com/therecipe



dep-tools:
#	go get -u github.com/ethereum/go-ethereum
#	go get -u github.com/therecipe/qt
#	go get -u github.com/sirupsen/logrus
#	go get -u github.com/dchest/siphash
	go install github.com/ethereum/go-ethereum/cmd/clef
	go install github.com/therecipe/qt
	go install github.com/therecipe/qt/cmd/qtdeploy
	go install github.com/therecipe/qt/cmd/qtsetup

install:
	go install ./cmd/clefui/main.go

build:
	go build -o ./build/clefui ./cmd/clefui/main.go

run: build
	./build/clefui

build-debug:
	go build -gcflags "all=-N -l" -o ./build/clefui ./cmd/clefui/main.go

debug: build-debug
	dlv --listen=:2345 --headless=true --api-version=2 exec ./build/clefui

deploy:
	qtdeploy build desktop ./cmd/clefui

.PHONY: build build-debug install deps setup dep-tools
