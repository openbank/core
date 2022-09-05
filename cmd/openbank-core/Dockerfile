FROM alpine:3.15

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

COPY cmd/openbank-core/core core

ENTRYPOINT ["sh", "-c", "./core"]
