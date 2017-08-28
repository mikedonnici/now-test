FROM golang:alpine
ADD . /go/src/github.com/mikedonnici/now-test
RUN go install github.com/mikedonnici/now-test
CMD ["/go/bin/now-test"]
EXPOSE 3000