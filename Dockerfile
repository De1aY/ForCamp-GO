FROM golang:latest
WORKDIR /usr/app
COPY . /usr/app

RUN go get -d -v ./...
RUN go build -o app
RUN chmod 777 ./app
RUN rm -rf /go/pkg