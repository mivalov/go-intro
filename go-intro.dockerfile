FROM golang:1.21.6-bookworm

RUN apt-get update && apt-get -y upgrade

RUN apt-get -y install \
    less \
    tree \
    vim

RUN apt-get -y autoremove \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

# RUN useradd -ms /bin/bash developer
# USER developer
WORKDIR /app

# Official Go language server
# https://pkg.go.dev/golang.org/x/tools/gopls
RUN go install golang.org/x/tools/gopls@latest
