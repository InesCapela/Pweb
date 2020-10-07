package main

import "fmt"
import "github.com/gin-gonic/gin"


func main() {
	fmt.Println("hello world")

	f := fib()
	fmt.Println(f(),f(),f(),f(),f())

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}

func fib() func() int{
	a, b :=0,1
	return func() int {
		a, b = b, a+b
		return a
	}
}