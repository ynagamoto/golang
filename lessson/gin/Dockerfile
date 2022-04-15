FROM golang:latest
COPY ./source /go/gin
WORKDIR /go/gin
RUN pwd
RUN ls -al
RUN go mod init example.com/gin-test
RUN go mod tidy
ENTRYPOINT go run main.go
