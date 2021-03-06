FROM golang:alpine as builder
RUN apk update && apk add git
COPY . $GOPATH/src/github.com/floge77/c2p/cloud2podcastdownloader
WORKDIR $GOPATH/src/github.com/floge77/c2p/cloud2podcastdownloader

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s'  -o /go/bin/cloud2podcastdownloader



#FROM python:3.7-alpine3.8
FROM alpine:3.8
RUN apk update && apk add curl python ffmpeg
RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl && chmod a+rx /usr/local/bin/youtube-dl
RUN adduser -D -g '' cloud2podcastdownloader
USER cloud2podcastdownloader

COPY --from=builder /go/bin/cloud2podcastdownloader /go/bin/cloud2podcastdownloader
VOLUME /downloads

ENTRYPOINT ["/go/bin/cloud2podcastdownloader"]
