FROM alpine:3.7

ARG VERSION
ARG USER=feynman
ARG GROUP=keen

LABEL "name"="random-names" \
      "version"="$VERSION" \
      "summary"="A random names generator application in Go." \
      "description"="Random names generates labels you can use in your applications or workflows." \
      "maintainer"="Robert Buck, buck.robert.j@gmail.com"

COPY random-names /usr/local/bin
COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh

RUN addgroup $GROUP && \
    adduser -S -G $GROUP $USER

RUN mkdir -p /random-names/config && \
    chown -R $USER:$GROUP /random-names

RUN set -eux \
    && apk add --no-cache ca-certificates curl dumb-init libcap su-exec \
    && random-names --version

EXPOSE 9000

ENTRYPOINT [ "docker-entrypoint.sh" ]
CMD [ "server", "--bind", "0.0.0.0" ]