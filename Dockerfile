FROM golang:1.18-alpine

RUN mkdir /app

ADD . /app

WORKDIR /app

# get git
RUN apk add git

RUN go get -d -v ./...

EXPOSE 9090

RUN go build -o main .

CMD ["./main"]