FROM alpine:3.15

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

COPY cmd/openbank-idp/idp idp

ENTRYPOINT ["sh", "-c", "./idp"]
