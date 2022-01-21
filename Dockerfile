FROM 10.0.45.243:21181/cloud-base/centos:7

ENV GIN_MODE=release

COPY config-example.yaml  /config.yaml
COPY helm-engine /

ENTRYPOINT [ "/helm-engine" ]
