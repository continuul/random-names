FROM alpine:3.7

LABEL maintainer="Robert Buck <bob@continuul.io>"

RUN addgroup fun && \
    adduser -S -G fun fun
RUN set -eux \
    && apk add --no-cache ca-certificates curl dumb-init libcap su-exec
#    && names-generator version

COPY random-names /usr/local/bin
COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
ENTRYPOINT [ "docker-entrypoint.sh" ]
CMD [ "generate" ]