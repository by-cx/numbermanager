FROM alpine:3.10.1 as builder

RUN apk update && apk upgrade
RUN apk add go git sqlite gcc g++ dep

ENV GOPATH /srv

ADD . /srv/src/github.com/by-cx/numbermanager
WORKDIR /srv/src/github.com/by-cx/numbermanager
RUN dep ensure
RUN go test
RUN go build -o numbermanager *.go


FROM alpine:3.10.1

EXPOSE 1323
WORKDIR /srv/app

COPY --from=builder /srv/src/github.com/by-cx/numbermanager/numbermanager /src/app/numbermanager

ENTRYPOINT ["/src/app/numbermanager"]