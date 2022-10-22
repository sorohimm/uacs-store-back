FROM debian:stable-slim
LABEL maintainer="sorohimm"

ENV DEBIAN_FRONTEND=noninteractive

COPY /build/uacs-store /usr/local/bin/uacs-store

# http gateway
EXPOSE 2604/tcp

# grpc api
EXPOSE 9000/tcp

WORKDIR /uacs

ENTRYPOINT ["/usr/local/bin/uacs-store"]