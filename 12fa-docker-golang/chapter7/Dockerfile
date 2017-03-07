FROM alpine:3.5

RUN apk add --no-cache curl

ADD build/ /app/

ENV PORT=3000

HEALTHCHECK --interval=60s --timeout=60s CMD curl -f http://localhost:$PORT/api/health || exit 1

WORKDIR /app

ENTRYPOINT ["/app/gotwitter"]