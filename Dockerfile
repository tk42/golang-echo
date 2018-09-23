FROM golang:1.11.0-stretch
MAINTAINER Tadashi KOJIMA <nsplat@gmail.com>

# install basic commands
RUN apt-get update && apt-get install -y vim less

# Install go dep
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep version
