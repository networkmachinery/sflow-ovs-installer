# Build Stage
FROM zanetworker/networkmachinery-build:1.12.4 AS build-stage

LABEL app="build-networkmachinery-sflow"
LABEL REPO="https://github.com/zanetworker/networkmachinery-sflow"

ENV PROJPATH=/go/src/github.com/zanetworker/networkmachinery-sflow

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/zanetworker/networkmachinery-sflow
WORKDIR /go/src/github.com/zanetworker/networkmachinery-sflow

RUN make build-alpine

# Final Stage
FROM zanetworker/networkmachinery-base:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/zanetworker/networkmachinery-sflow"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/networkmachinery-sflow/bin

WORKDIR /opt/networkmachinery-sflow/bin

COPY --from=build-stage /go/src/github.com/zanetworker/networkmachinery-sflow/bin/networkmachinery-sflow /opt/networkmachinery-sflow/bin/
RUN chmod +x /opt/networkmachinery-sflow/bin/networkmachinery-sflow

# Create appuser
RUN adduser -D -g '' networkmachinery-sflow
USER networkmachinery-sflow

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/networkmachinery-sflow/bin/networkmachinery-sflow"]
