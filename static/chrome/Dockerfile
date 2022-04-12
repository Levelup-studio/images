ARG VERSION
FROM --platform=linux/arm64 golang:alpine3.15 as go
ENV PATH="${PATH}:/sbin"
COPY devtools /devtools
RUN \
    apk update && \
    apk --no-cache add libx11-dev build-base && \
    wget https://dl-cdn.alpinelinux.org/alpine/v3.11/community/aarch64/upx-3.95-r2.apk && \
    apk add --allow-untrusted upx-3.95-r2.apk && \
    cd /devtools && \
    go test -race && \
    GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" && \
    upx /devtools/devtools

FROM --platform=linux/arm64 selenoid/dev_chrome:$VERSION
ENV PATH="/usr/bin/chromedriver:${PATH}"
ENV PATH="${PATH}:/sbin"
ENV DBUS_SESSION_BUS_ADDRESS=/dev/null
COPY --from=go /devtools/devtools /usr/bin/
COPY entrypoint.sh /
RUN chmod +x /usr/bin/chromedriver
USER selenium

EXPOSE 4444
ENTRYPOINT ["/entrypoint.sh"]
