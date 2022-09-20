FROM golang:1.19

WORKDIR /go/src/app
COPY . .

RUN go get
RUN go build -o pricing-engine

EXPOSE 8080

ENTRYPOINT ./pricing-engine