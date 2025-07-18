
FROM golang:1.24-alpine

WORKDIR /app

COPY  . .

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GO111MODULE=on

RUN go mod tidy

RUN go build -o app .


CMD ["./app"]

