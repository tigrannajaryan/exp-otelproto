FROM alpine:3.9

RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN apk add --update ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
ADD server /app/

CMD ["/app/server"]

