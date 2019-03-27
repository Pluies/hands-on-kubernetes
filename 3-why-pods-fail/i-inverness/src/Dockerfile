FROM golang:1.12-alpine as builder

ADD main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o main .

# Make things even smaller!
RUN apk add --no-cache upx \
 && upx main

FROM scratch
COPY --from=builder /go/main /
CMD ["/main"]
