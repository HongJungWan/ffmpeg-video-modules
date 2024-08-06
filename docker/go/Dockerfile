FROM golang:1.20-alpine

WORKDIR /go/src/github.com/HongJungWan/ffmpeg-video-modules/app

RUN apk add --no-cache git

COPY ./app/go.mod ./app/go.sum ./

RUN go mod download

COPY ./app .

RUN go build -o /main .

CMD ["/main"]