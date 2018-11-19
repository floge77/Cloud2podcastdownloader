FROM golang:alpine
RUN apk update && apk add curl git python ffmpeg
RUN adduser -D -g '' cloud2podcastdownloader
RUN curl -L https://yt-dl.org/downloads/latest/youtube-dl -o /usr/local/bin/youtube-dl && chmod a+rx /usr/local/bin/youtube-dl
COPY . $GOPATH/src/cloud2podcastdownloader
WORKDIR $GOPATH/src/cloud2podcastdownloader

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags='-w -s'  -o /go/bin/cloud2podcastdownloader
ENV PATH=/usr/local/python:/usr/local/bin:$PATH
VOLUME /downloads

ENTRYPOINT ["/go/bin/cloud2podcastdownloader"]