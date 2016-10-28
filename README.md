# k8s-example

Simple Echo Service implemented in Go.  Deployed using Docker and Kubernetes

#### To run:

##### Execute the command

```shell

$ go get github.com/karlmoad/k8s-example

```

navigate to the github.com/karlmoad/k8s-example directory within your local `$GOPATH`

##### Execute the build

```shell
$ docker build -t k8s-example:latest .
```

##### Run the image

```shell
$ docker run -p 8080:8080 -d k8s-example:latest

```