
# Start from golang base image
FROM golang:alpine as builder

# Add Maintainer info
LABEL maintainer="Wahome <kevowahome@gmail.com>"

# Enable go modules
ENV GO111MODULE=on
ENV GOOS="linux"
ENV CGO_ENABLED=0
ENV GOARCH=amd64

# Install git. (alpine image does not have git in it)
RUN apk update && \
    apk add curl \
            git \
            bash \
            make \
            tar \
            ca-certificates && \
    rm -rf /var/cache/apk/*

# Set current working directory
WORKDIR /app

# Note here: To avoid downloading dependencies every time we
# build image. Here, we are caching all the dependencies by
# first copying go.mod and go.sum files and downloading them,
# to be used every time we build the image if the dependencies
# are not changed.

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download
RUN go mod verify

# Now, copy the source code
COPY . ./

# Note here: CGO_ENABLED is disabled for cross system compilation
# It is also a common best practise.

# Build the application.
RUN go build -v -o /bin/main

RUN mkdir -p /var/log/app

EXPOSE 8080
