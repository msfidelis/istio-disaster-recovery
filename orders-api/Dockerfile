FROM golang:1.15 AS builder

WORKDIR $GOPATH/src/orders-api

COPY . ./

RUN go get -u
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:3.12.3

COPY --from=builder /go/src/orders-api/main ./
COPY --from=builder /go/src/orders-api/configs ./configs

EXPOSE 8080

ENTRYPOINT ["./main"]
