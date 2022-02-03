FROM BASEIMAGE
RUN apk --no-cache add ca-certificates bash curl

RUN curl -sL --http1.1 https://cnfl.io/cli | sh -s -- latest

ARG ARCH
ARG TINI_VERSION

ADD provider /usr/local/bin/crossplane-confluent-provider

EXPOSE 8080
USER 1001
ENTRYPOINT ["crossplane-confluent-provider"]
