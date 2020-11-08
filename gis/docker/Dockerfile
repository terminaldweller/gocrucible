FROM alpine:3.12
RUN apk update && apk upgrade
RUN apk add git go

RUN go get github.com/go-kit/kit/transport/http
RUN go get github.com/go-kit/kit/endpoint
RUN go get github.com/go-kit/kit/endpoint

RUN git clone http://github.com/terminaldweller/gocrucible
EXPOSE 8080
RUN (cd gocrucible/gis && go build  && ln /gocrucible/gis/gis /usr/bin/gis)
