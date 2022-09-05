FROM node:17-alpine

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

COPY docs /root

ENTRYPOINT ["sh", "-c", "node build"]
