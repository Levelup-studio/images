# This is fork with Apple M1 support
Ubuntu was replaced with alpine for better arm64 support. At the moment it  supports only chomium browser (instead of google chrome).
# How to build own image
* install GO (you can use brew since you're using MacOS)
* clone this repo
* in root repo directory run
```
go install github.com/markbates/pkger/cmd/pkger
go generate github.com/aerokube/images
go build
```
* go to `images/selenium/base` and run `docker build -t hubhead:selenoid-base-arm .`
* go back to root repo directory and run `./images chrome -b 100.0.4896.60-1 -d latest -t selenoid-arm/chrome:100` (you can specify any image tag after -t)
* now you have image `selenoid-arm/chrome:100` that runs chromium and works ok with Apple M1

Only 99.0.4844.84-r0 and 100.0.4896.60-1 chromium versions are valid. Somehow older versions are not working on alpine 3.15.


# Browser Images
[![Build Status](https://github.com/aerokube/images/workflows/build/badge.svg)](https://github.com/aerokube/images/actions?query=workflow%3Abuild)
[![Release](https://img.shields.io/github/release/aerokube/images.svg)](https://github.com/aerokube/images/releases/latest)

This repository contains [Docker](http://docker.com/) build files to be used for [Selenoid](http://github.com/aerokube/selenoid) and [Moon](http://github.com/aerokube/moon) projects. You can find prebuilt images [here](https://hub.docker.com/u/selenoid/).

## Download Statistics

### Firefox: [![Firefox Docker Pulls](https://img.shields.io/docker/pulls/selenoid/firefox.svg)](https://hub.docker.com/r/selenoid/firefox)

### Chrome: [![Chrome Docker Pulls](https://img.shields.io/docker/pulls/selenoid/chrome.svg)](https://hub.docker.com/r/selenoid/chrome)

### Opera: [![Opera Docker Pulls](https://img.shields.io/docker/pulls/selenoid/opera.svg)](https://hub.docker.com/r/selenoid/opera)

### Android: [![Android Docker Pulls](https://img.shields.io/docker/pulls/selenoid/android.svg)](https://hub.docker.com/r/selenoid/android)

## Building Images

Moved to: http://aerokube.com/images/latest/#_building_images

## Image information
Moved to: http://aerokube.com/images/latest/#_browser_image_information
