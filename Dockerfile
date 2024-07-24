FROM alpine:3.20

WORKDIR /app

COPY ./gateway /app

CMD ["/app/gateway"]
