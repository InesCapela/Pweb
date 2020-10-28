// docker run --name log_bd-gin-api -p3306:3306 -e MYSQL_ROOT_PASSWORD=12345678 -e MYSQL_DATABASE=demo -d log_bd:5.7
//docker build -t name .
//
package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type (
	// todoModel describes a todoModel type
	todoModel struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed"`
	}

	// transformedTodo represents a formatted todo
	transformedTodo struct {
		ID        uint      `json:"id"`
		Title     string    `json:"title"`
		Completed bool      `json:"completed"`
		Time      time.Time `json:"time"`
	}
)

var db *gorm.DB

func init() {
	// http://gorm.io/docs/connecting_to_the_database.html
	//open a db connection
	var err error
	db, err = gorm.Open("mysql", "root:12345678@tcp(mysql:3306)/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	db.AutoMigrate(&todoModel{})
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1/todos")
	{
		v1.GET("/", fetchAllTodo)
		v1.GET("/:id", fetchSingleTodo)
		v1.POST("/", createTodo)
		v1.PUT("/:id", updateTodo)
		v1.DELETE("/id/:id", deleteTodo)
		v1.DELETE("/all", deleteAllTodo)
	}
	router.Run()
}

// createTodo add a new todo
func createTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := todoModel{Title: c.PostForm("title"), Completed: completed}

	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": todo.ID})
}

// fetchAllTodo fetch all todos
func fetchAllTodo(c *gin.Context) {
	var todos []todoModel
	var _todos []transformedTodo

	db.Find(&todos)

	if len(todos) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusOK, "data": "Not Found"})
		return
	}
	for i, todo := range todos {
		_todos[i] = transformedTodo{ID: todo.ID, Title: todo.Title, Completed: todo.Completed == 1}
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todos})
}

// fetchSingleTodo fetch a single todo
func fetchSingleTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	db.Find(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusOK, "data": "Not Found"})
		return
	}
	_todo := transformedTodo{ID: todo.ID, Title: todo.Title, Completed: todo.Completed == 1}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _todo})
}

// updateTodo update a todo
func updateTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	db.Find(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusOK, "data": "Not Found"})
		return
	}
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo.Title = c.PostForm("title")
	todo.Completed = completed

	db.Save(&todo)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo updated successfully!"})
}

// deleteTodo remove a todo
func deleteTodo(c *gin.Context) {
	var todo todoModel
	todoID := c.Param("id")

	db.Find(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusOK, "data": "Not Found"})
		return
	}

	db.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Todo deleted successfully!"})
}

func deleteAllTodo(c *gin.Context) {
	var todo []todoModel

	db.Find(&todo)
	db.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "All records have been deleted successfully!"})
}
