FROM golang:1.13

WORKDIR /go/src/astral-test
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 3000

CMD ["astral-test"]
