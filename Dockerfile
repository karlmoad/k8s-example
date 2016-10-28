FROM scratch
MAINTAINER Karl Moad <https:/github.com/karlmoad>
ADD k8s-example k8s-example
ENTRYPOINT [/k8s-example]
EXPOSE 8080