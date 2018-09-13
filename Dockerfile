FROM golang:1.11.0-stretch
MAINTAINER Tadashi KOJIMA

# install basic commands
RUN apt-get install -y vim less

# Install go dep
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep version
