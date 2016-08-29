FROM scratch

ADD https://trilliumstaffing.com/gh-latest/repo/TrilliumIT/gh-latest/gh-latest /

EXPOSE 8080

CMD ["/gh-latest"]
