FROM BASEIMAGE
RUN apk --no-cache add ca-certificates bash curl

RUN curl -sL --http1.1 https://cnfl.io/cli | sh -s -- -b /usr/local/bin latest

ARG ARCH
ARG TINI_VERSION

ADD provider /usr/local/bin/crossplane-confluent-provider

RUN adduser -D -u 2000 -h /home/appuser appuser
RUN mkdir /.confluent && chown appuser:appuser /.confluent

WORKDIR /home/appuser

EXPOSE 8080
USER appuser
ENTRYPOINT ["crossplane-confluent-provider"]
