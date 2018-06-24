FROM golang:1.8.5-jessie

ENV PORT 3000

WORKDIR /go/src/github.com/grzegorz-bielski/microstream/backend

RUN go get github.com/golang/dep/cmd/dep

ADD . .
RUN dep ensure --vendor-only
RUN go generate
RUN go install -v ./...

CMD ["backend"]