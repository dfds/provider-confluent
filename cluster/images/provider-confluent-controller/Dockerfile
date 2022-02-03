FROM BASEIMAGE
RUN apk --no-cache add ca-certificates bash curl

RUN curl -sL --http1.1 https://cnfl.io/cli | sh -s -- latest

ARG ARCH
ARG TINI_VERSION

ADD provider /usr/local/bin/crossplane-confluent-provider

RUN adduser -D -u 1001 -h /home/appuser appuser
RUN mkdir /.confluent && chown 1001:1001 /.confluent

WORKDIR /home/appuser

EXPOSE 8080
USER appuser
ENTRYPOINT ["crossplane-confluent-provider"]
