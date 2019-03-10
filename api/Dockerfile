FROM golang:alpine

WORKDIR /go/src/api
COPY . .

RUN go install -v ./...

CMD ["api"]
