FROM golang:alpine as builder

ENV GOPROXY https://goproxy.cn/

WORKDIR /go/release
RUN apk update && apk add tzdata

COPY go.mod ./go.mod
RUN go mod tidy
COPY . .
RUN pwd && ls

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o webapp .

FROM alpine

COPY --from=builder /go/release/webapp /

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

EXPOSE 8080

CMD ["/webapp"]
