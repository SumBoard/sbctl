SBCTL_VERSION?=dev
.PHONY: build
build:
	CGO_ENABLED=0 go build -trimpath -o build/ -ldflags "-w -X github.com/SumBoard/sbctl/internal/build.Version=$(SBCTL_VERSION)" .

.PHONY: clean
clean:
	rm -rf build/

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./...

.PHONY: vet
vet:
	go vet ./...
