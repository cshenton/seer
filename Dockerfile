FROM golang:latest AS build-env

ENV GOPATH=/go
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0

ADD . /go/src/github.com/cshenton/seer
RUN cd /go/src/github.com/cshenton/seer && \
    go get ./... && \
    go build -a -v -o ./runseer ./main.go


FROM alpine
WORKDIR /
COPY --from=build-env /go/src/github.com/cshenton/seer/runseer /seer
ENTRYPOINT ./seer