FROM docker.io/library/golang:1 AS builder
WORKDIR /usr/src/app
COPY . ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -a -installsuffix cgo -o bin/s3browser

FROM docker.io/library/alpine:latest
WORKDIR /usr/src/app
RUN addgroup -S s3browser && adduser -S s3browser -G s3browser
RUN apk add --no-cache \
  ca-certificates \
  dumb-init
COPY --from=builder --chown=s3browser:s3browser /usr/src/app/bin/s3browser ./
USER s3browser
EXPOSE 8080
ENTRYPOINT [ "/usr/bin/dumb-init", "--" ]
CMD [ "/usr/src/app/s3browser" ]
