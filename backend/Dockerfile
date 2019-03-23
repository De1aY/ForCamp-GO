FROM golang:1.12.1

WORKDIR $GOPATH/src/wplay
COPY . .

RUN go get
EXPOSE 80 443

CMD ["go", "run", "main.go"]
