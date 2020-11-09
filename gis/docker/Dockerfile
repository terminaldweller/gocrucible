FROM alpine:3.12 as builder
RUN apk update && apk upgrade
RUN apk add git go

RUN go get github.com/go-kit/kit/transport/http
RUN go get github.com/go-kit/kit/endpoint
RUN go get github.com/go-kit/kit/endpoint

RUN git clone http://github.com/terminaldweller/gocrucible
RUN (cd gocrucible/gis && go build)

FROM alpine:3.12
WORKDIR /root/
COPY --from=builder /gocrucible/gis/gis .
EXPOSE 8080
