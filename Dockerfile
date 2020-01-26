FROM golang:1.13.6-alpine3.11

WORKDIR ./app

COPY . .

RUN go build .

CMD ["./titan"]