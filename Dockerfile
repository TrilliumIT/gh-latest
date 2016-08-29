FROM scratch

ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/certs/ca-certificates.crt
ADD https://trilliumstaffing.com/gh-latest/repo/TrilliumIT/gh-latest/gh-latest /

CMD ["/gh-latest"]
