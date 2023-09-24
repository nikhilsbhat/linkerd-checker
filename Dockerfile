### Description: Dockerfile for gocd-prometheus-exporter
FROM alpine:3.18

COPY linkerd-checker /

# Starting
ENTRYPOINT [ "/linkerd-checker" ]