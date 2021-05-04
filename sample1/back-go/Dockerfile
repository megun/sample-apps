FROM golang:1.16.3

WORKDIR /go/src/app

COPY main.go main.go

ENV GO111MODULE=off
RUN go get github.com/go-sql-driver/mysql

CMD ["go", "run", "main.go"]
