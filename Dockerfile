FROM golang
ADD . /go/src/github.com/laszlovaspal/devops-challenge

RUN GO15VENDOREXPERIMENT=1 go install github.com/laszlovaspal/devops-challenge

ENTRYPOINT ["/go/bin/devops-challenge"]
CMD ["-help"]

EXPOSE 8080
