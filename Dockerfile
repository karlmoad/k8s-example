FROM golang:latest
RUN mkdir -p /go/src/github.com/karlmoad/k8s-example
ADD . /go/src/github.com/karlmoad/k8s-example
RUN go get github.com/gorilla/mux
RUN go install github.com/karlmoad/k8s-example
ENTRYPOINT /go/bin/k8s-example
EXPOSE 8080