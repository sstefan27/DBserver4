FROM golang:1.16.4-stretch as builder
ENV GOPATH /go
WORKDIR /
COPY ./ /

RUN go build main.go

FROM debian:stretch
COPY --from=0 /main .
COPY --from=0 /.env /.env
COPY --from=0 /db/migrate /db/migrate
ENTRYPOINT ["./main"]