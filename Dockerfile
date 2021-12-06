FROM golang:alpine

MAINTAINER mehmet okdemir

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /go/src/app
COPY . .

RUN go get -d -v
RUN go build -v

CMD ["./cryptocurrency-portfolio"]

