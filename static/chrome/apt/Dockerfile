FROM --platform=linux/arm64 hubhead:selenoid-base-arm

ARG VERSION
ARG PACKAGE=chromium-browser
ARG INSTALL_DIR=chromium-browser

LABEL browser=$PACKAGE:$VERSION
ENV PATH="${PATH}:/sbin"
RUN \
        apk add chromium=$VERSION chromium-chromedriver=$VERSION && \
        ( [ "$PACKAGE" != "chromium-browser" ] && ln /usr/bin/$PACKAGE /usr/bin/chromium-browser ) || true && \
        chromium-browser --version
