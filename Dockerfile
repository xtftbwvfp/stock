FROM alpine

WORKDIR /srv/stock

ADD ./dist/stock.tar.gz /srv/

EXPOSE 4869 4870
ENTRYPOINT ["./stock", "-c", "./config.toml"]
