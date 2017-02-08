FROM alpine:3.5

RUN apk add --no-cache curl

ADD build/ /app/

HEALTHCHECK --interval=5s --timeout=2s CMD PORT=${PORT:-3000} curl -f http://localhost:$PORT/api/health || exit 1

WORKDIR /app

ENTRYPOINT ["/app/gotwitter"]