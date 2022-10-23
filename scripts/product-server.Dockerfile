FROM debian:stable-slim
LABEL maintainer="sorohimm"

ENV DEBIAN_FRONTEND=noninteractive

COPY /build/product /usr/local/bin/product

# http gateway
EXPOSE 2604/tcp

# grpc api
EXPOSE 9000/tcp

ENTRYPOINT ["/usr/local/bin/product"]