FROM debian:stable-slim
LABEL maintainer="sorohimm"

ENV DEBIAN_FRONTEND=noninteractive

COPY /build/uacs-auth /usr/local/bin/uacs-auth

# http gateway
EXPOSE 2104/tcp

# grpc api
EXPOSE 9001/tcp

WORKDIR /uacs

ENTRYPOINT ["/usr/local/bin/uacs-auth"]