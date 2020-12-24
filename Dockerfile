FROM alpine:3.12.1
ARG VERSION
COPY bin/swage /usr/bin/swage
ENTRYPOINT ["/usr/bin/swage"]
