FROM golang:1.14

WORKDIR /go/src/github.com/aankittcoolest/bookstore_users-api
COPY . .

RUN go get github.com/codegangsta/gin
RUN go get -d ./...

CMD ["gin", "run", "main.go"]