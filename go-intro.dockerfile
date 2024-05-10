FROM golang:1.22.3-bookworm

RUN apt-get update \
    && apt-get -y upgrade \
    && apt-get -y install \
        less \
        tree \
        vim \
    && apt-get -y autoremove \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

RUN useradd -ms /bin/bash gopher
USER gopher
WORKDIR /home/gopher/go-intro

# Official Go language server
# https://pkg.go.dev/golang.org/x/tools/gopls
RUN go install golang.org/x/tools/gopls@latest
