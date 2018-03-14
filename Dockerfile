FROM golang:alpine AS build-env
ADD . /src
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
RUN cd /src && go build -a -v -o ./seer ./main.go

FROM alpine
WORKDIR /app
COPY --from=build-env /src/seer /app/seer
ENTRYPOINT ./seer