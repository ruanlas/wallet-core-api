FROM alpine:3

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

USER nonroot

WORKDIR /wallet
COPY ./build/app ./

CMD ["/wallet/app"]