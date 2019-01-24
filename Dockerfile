FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

ADD https://trilliumstaffing.com/gh-latest/repo/TrilliumIT/gh-latest/gh-latest /
RUN chmod +x gh-latest

EXPOSE 8080

CMD ["/gh-latest"]
