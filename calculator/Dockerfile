FROM golang:latest

RUN go get github.com/gin-gonic/gin

COPY . .

LABEL maintainer="Ines <36502@ufp.edu.pt>" \
      version="1.0"

CMD go run main.go