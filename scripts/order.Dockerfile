FROM debian:stable-slim
LABEL maintainer="sorohimm"

ENV DEBIAN_FRONTEND=noninteractive

COPY /build/order /usr/local/bin/order

# http gateway
EXPOSE 2104/tcp

# grpc api
EXPOSE 9001/tcp

ENTRYPOINT ["/usr/local/bin/order"]