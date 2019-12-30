# FROM golang as build

# ARG COREDNS_VERSION=1.6.4

# WORKDIR /tmp

# COPY plugin.cfg /tmp/plugin.cfg

# RUN set -x \
#   && apt update && apt install curl unzip make -y \
#   && curl -Lo coredns.zip https://github.com/coredns/coredns/archive/v$COREDNS_VERSION.zip \
#   && unzip coredns.zip \
#   && cd coredns-$COREDNS_VERSION \
#   && cp /tmp/plugin.cfg . \
#   && make \
#   && ./coredns -plugins \
#   && mv coredns /tmp

FROM busybox

# get the config file from the local folder
COPY Corefile /etc/coredns/Corefile

# get thecoredns binary from internal coredn with plugins existing image
# COPY coredns /opt/coredns
COPY --from=registry:5000/infra/corednswithplugins:1.6.4-drop /tmp/coredns /opt/coredns
# COPY --from=build /tmp/coredns /opt/coredns

EXPOSE 53/udp 15353 19253

CMD ["/opt/coredns", "-conf", "/etc/coredns/Corefile"]
