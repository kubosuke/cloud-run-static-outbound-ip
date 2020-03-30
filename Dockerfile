FROM golang:alpine as builder

WORKDIR /go/app

COPY . .

RUN apk add --no-cache git && go build -o app

FROM alpine

WORKDIR /app

COPY --from=builder /go/app/. .

RUN addgroup go && adduser -D -G go go && chown -R go:go /app/app

RUN apk add --no-cache ca-certificates bash openssh-client

ENV TINI_VERSION v0.18.0
ENV HTTP_PROXY localhost:5000

ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini-static /tini

RUN chmod +x /tini
RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["/tini", "--", "./entrypoint.sh"]
