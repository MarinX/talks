#!/bin/bash

# Download source code

go get -d github.com/tensorflow/tensorflow/tensorflow/go

# Build the lib

cd ${GOPATH}/src/github.com/tensorflow/tensorflow
./configure
bazel build --config opt //tensorflow:libtensorflow.so

# Make lib available to linker

sudo cp ${GOPATH}/src/github.com/tensorflow/tensorflow/bazel-bin/tensorflow/libtensorflow.so /usr/local/lib
