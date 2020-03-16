GO_LDFLAGS := -extldflags "-static"
GO_LDFLAGS += -w -s

.PHONY: build
build: clean
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -ldflags '$(GO_LDFLAGS)' -o hello-pods

.PHONY: clean
clean:
	rm -f hello-pods

.PHONY: image
image: build
	docker build -t registry.gitlab.com/grzesiek/hello-pods:latest .
