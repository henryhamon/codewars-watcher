FROM golang:latest

RUN mkdir -p /go/src \
&& mkdir -p /go/bin \
&& mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/src/github.com/leometzger/codewars-watcher
ADD . $GOPATH/src/github.com/leometzger/codewars-watcher

WORKDIR $GOPATH/src/github.com/leometzger/codewars-watcher

RUN go get && go build .

EXPOSE 8080
