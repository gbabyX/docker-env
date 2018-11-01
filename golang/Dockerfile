FROM golang:latest

RUN mkdir -p /go/src/app
COPY app.go /go/src/app
WORKDIR /go/src/app
RUN go get github.com/gin-gonic/gin
RUN go build /go/src/app/app.go
EXPOSE 80
CMD ["./app"]