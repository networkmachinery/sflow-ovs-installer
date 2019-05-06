# Build Stage
FROM golang:1.12.4 AS build-stage

LABEL app="build-networkmachinery-sflow"
LABEL REPO="https://github.com/networkmachinery/sflow-ovs-installer"

ENV PROJPATH=/go/src/github.com/networkmachinery/sflow-ovs-installer

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/networkmachinery/sflow-ovs-installer
WORKDIR /go/src/github.com/networkmachinery/sflow-ovs-installer

RUN make build-alpine

# Final Stage
FROM ubuntu:16.04

RUN apt update && apt install -y openvswitch-switch iptables wget
# Adding dump-init to handle
RUN wget https://github.com/Yelp/dumb-init/releases/download/v1.2.2/dumb-init_1.2.2_amd64.deb
RUN dpkg -i dumb-init_*.deb

ARG GIT_COMMIT
ARG VERSION

LABEL REPO="https://github.com/networkmachinery/sflow-ovs-installer"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/sflow-ovs-installer/bin

WORKDIR /opt/sflow-ovs-installer/bin

COPY --from=build-stage /go/src/github.com/networkmachinery/sflow-ovs-installer/bin/sflow-ovs-installer /opt/sflow-ovs-installer/bin/
RUN chmod +x /opt/sflow-ovs-installer/bin/sflow-ovs-installer

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["/opt/sflow-ovs-installer/bin/sflow-ovs-installer"]
