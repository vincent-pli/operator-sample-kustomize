FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/operator-sample-kustomize \
    USER_UID=1001 \
    USER_NAME=operator-sample-kustomize

# install templates
COPY templates /usr/local/templates
# install operator binary
COPY build/_output/bin/operator-sample-kustomize ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
