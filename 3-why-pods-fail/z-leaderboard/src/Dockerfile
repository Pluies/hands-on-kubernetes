FROM golang:1.12

WORKDIR /go/src/github.com/pluies/hands-on-kubernetes
ENV GO111MODULE=on
ADD go.mod .
ADD go.sum .
RUN go mod download

ADD *.go .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .
ADD templates templates/
CMD ["./main"]
