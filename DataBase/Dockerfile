FROM golang:latest

RUN go get github.com/gin-gonic/gin
RUN go get github.com/jinzhu/gorm
RUN go get github.com/jinzhu/gorm/dialects/mysql

COPY . .

LABEL maintainer="Ines <36502@ufp.edu.pt>" \
      version="1.0"

CMD go run main_todo.go