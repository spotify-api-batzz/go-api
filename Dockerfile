FROM golang:1.16.3-alpine

COPY src /test
WORKDIR /test

RUN env GOOS=linux GOARCH=amd64 go build
RUN chmod +x /test/spotify-server

EXPOSE 8080

CMD [ "/test/spotify-server" ]  