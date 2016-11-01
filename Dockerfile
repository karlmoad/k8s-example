FROM scratch
MAINTAINER Karl Moad <https:/github.com/karlmoad>
ADD k8s-example /
EXPOSE 8080
CMD ["/k8s-example"]
