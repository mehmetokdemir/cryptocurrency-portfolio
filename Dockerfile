FROM golang:1.17.1

WORKDIR /go/src/app
COPY . .

RUN go get -d -v
RUN go build -v

# CMD
CMD ["./cryptocurrency-portfolio"]