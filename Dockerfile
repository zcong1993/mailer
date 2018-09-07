FROM zcong/golang:1.10.3 AS build
WORKDIR /go/src/github.com/zcong1993/mailer
COPY . .
RUN dep ensure -vendor-only -v && \
    CGO_ENABLED=0 go build -o ./bin/mailer main.go

FROM alpine:3.7
WORKDIR /opt
RUN apk add --no-cache ca-certificates
COPY --from=build /go/src/github.com/zcong1993/mailer/bin/* /usr/bin/
CMD ["mailer"]
