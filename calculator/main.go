//docker run -it --rm -p "8080:8080" -v "C:\Users\maria\go\src\calculator\:/app" go
package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	gin.ForceConsoleColor() // enable color console output
	gin.Logger()            // allow fmt.print to output

	r := gin.Default()
	r.LoadHTMLGlob("index.html")

	g := r.Group("/")
	{

		g.GET("/add", form("/add", "Add numbers"))
		g.POST("/add", add)

		g.GET("/sub", form("/sub", "Subtract numbers"))
		g.POST("/sub", sub)

		g.GET("/div", form("/div", "Divide numbers"))
		g.POST("/div", div)

		g.GET("/mul", form("/mul", "Multiply numbers"))
		g.POST("/mul", mul)
	}

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


func form(formAction, buttonText string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"formAction": formAction,
			"buttonText": buttonText,
		})
	}
}

/*
 *	Get and validate numbers from Post Form
 */
func getAndValidateNumbers(c *gin.Context) (int, int, error) {

	stringA, foundA := c.GetPostForm("a")
	stringB, foundB := c.GetPostForm("b")

	//error
	if !foundA || !foundB {
		err := fmt.Errorf("invalid numbers")
		return 0, 0, err
	}

	a, _ := strconv.Atoi(stringA)
	b, _ := strconv.Atoi(stringB)

	fmt.Println("A = ", a)
	fmt.Println("B = ", b)
	fmt.Println()

	return a, b, nil
}

func add(c *gin.Context) {

	a, b, err := getAndValidateNumbers(c)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": a + b,
	})
}

func sub(c *gin.Context) {
	a, b, err := getAndValidateNumbers(c)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": a - b,
	})
}

func div(c *gin.Context) {
	a, b, err := getAndValidateNumbers(c)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": a / b,
	})
}

func mul(c *gin.Context) {
	a, b, err := getAndValidateNumbers(c)

	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": a * b,
	})
}
