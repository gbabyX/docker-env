FROM golang:latest

ARG PROJECT_DIR=/var/www
RUN mkdir -p ${PROJECT_DIR}
WORKDIR ${PROJECT_DIR}
RUN go get -u github.com/gin-gonic/gin
ENTRYPOINT [ "go","run","/var/www/main.go"]
EXPOSE 8080