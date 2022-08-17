FROM golang:1.19

MAINTAINER "Joshua Schmitt <jqshuv@gmail.com>"

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./main.go

CMD ["app"]