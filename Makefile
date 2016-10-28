.PHONY: all build container push clean

TAG ?= 1.0.0
PREFIX ?= karlmoad

all: container

build: service.go
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o k8s-example --ldflags '-w' ./service.go

container: build
	docker build -t $(PREFIX)/k8s-example:$(TAG) .

push:
	docker push $(PREFIX)/k8s-example:$(TAG)

clean:
	rm -f k8s-example
