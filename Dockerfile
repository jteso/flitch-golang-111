FROM alpine:3.1
MAINTAINER Flitch
ADD release/hello-v1.0.0-linux-amd64 /usr/bin/app
ENTRYPOINT ["app"]