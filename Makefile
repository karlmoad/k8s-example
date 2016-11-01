.PHONY: all build container clean

TAG ?= 1.0.0
PREFIX ?= moadkj

all: container

build: service.go
	GOOS=linux CGO_ENABLED=0 go build -a -installsuffix cgo -o k8s-example ./service.go

container: build
	docker build -t $(PREFIX)/k8s-example:$(TAG) .

clean:
	rm -f k8s-example
