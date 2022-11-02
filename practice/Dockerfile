# Multi-stage image
# Why need of it ?
# Because once we build app, we don't require go to be installed. 
# To save image size, remove go.
# That where comes the concept of builder.
# Earlier images size was 354MB, and now 126MB


FROM golang:alpine3.15 as builder

WORKDIR /path

COPY . .

RUN go build -o "app"


FROM alpine:3.15.5

WORKDIR /path

RUN apk update \
  && apk -U upgrade \
  && apk add --no-cache ca-certificates bash gcc \
  && update-ca-certificates --fresh \
  && rm -rf /var/cache/apk/*

COPY --from=builder /path/app  .

RUN chmod +x /path/app

ENTRYPOINT ["/path/app"]
