FROM --platform=linux/arm64 golang:alpine3.15 as go

COPY xseld /xseld

COPY fileserver /fileserver

RUN \
    apk update && \
    apk --no-cache add libx11-dev build-base && \
    wget https://dl-cdn.alpinelinux.org/alpine/v3.11/community/aarch64/upx-3.95-r2.apk && \
    apk add --allow-untrusted upx-3.95-r2.apk && \
    cd /xseld && \
    go env -w GO111MODULE=auto &&\
    if [ `uname -m` = "aarch64" ] ; then \
       GOOS=linux GOARCH=arm64 go build -ldflags="-s -w"; \
    else \
       GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"; \
    fi && \
    upx /xseld/xseld && \
    cd /fileserver && \
    go test -race && \
    if [ `uname -m` = "aarch64" ] ; then \
       GOOS=linux GOARCH=arm64 go build -ldflags="-s -w"; \
    else \
       GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"; \
    fi && \
    upx /fileserver/fileserver

FROM --platform=linux/arm64 alpine:3.15
RUN \
   echo 'https://dl-cdn.alpinelinux.org/alpine/edge/main/' >> /etc/apk/repositories &&\
   echo 'https://dl-cdn.alpinelinux.org/alpine/edge/community/' >> /etc/apk/repositories &&\
   apk update

RUN \
    apk add xvfb-run xorg-server xsel x11vnc feh fluxbox xvfb pulseaudio-utils nano &&\
    apk add wmctrl --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing/ &&\
    apk add pulseaudio --repository=http://dl-cdn.alpinelinux.org/alpine/edge/community/

RUN \
    adduser --system --home /home/selenium --uid 4096 \
    --ingroup root --disabled-password --shell /bin/bash selenium && \
    mkdir -p /home/selenium/Downloads && \
    mkdir -p /home/selenium/.fluxbox && \
    chgrp -R 0 /home/selenium && \
    chmod -R g=u /home/selenium && \
    rm -Rf /tmp/* && rm -Rf /var/lib/apt/lists/*

COPY fluxbox /usr/share/fluxbox/styles/
COPY --chown=selenium:root fluxbox /home/selenium/.fluxbox/
COPY aerokube.png /usr/share/images/fluxbox/

COPY --from=go /fileserver/fileserver /usr/bin/
COPY --from=go /xseld/xseld /usr/bin/
