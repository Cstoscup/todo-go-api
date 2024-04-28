package main

import (
	"net/http"

	"example/todo-go-api/database"
	"example/todo-go-api/model"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type todo struct {
	ID        string `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Task: "Task 1", Completed: false},
	{ID: "2", Task: "Task 2", Completed: false},
	{ID: "3", Task: "Task 3", Completed: false},
}

func main() {
	loadEnv()
	loadDatabase()

	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", createTodo)
	router.PUT("/todos/:id", updateTodo)
	// router.DELETE("/todos/:id", deleteTodo)
	router.Run("localhost:8080")
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.Connect()
	database.DB.AutoMigrate(&model.Todo{})
}

func getTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func createTodo(c *gin.Context) {
	var newTodo todo

	err := c.BindJSON(&newTodo)

	if err != nil {
		return
	}

	todos = append(todos, newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func updateTodo(c *gin.Context) {
	id := c.Param("id")
	var updatedTodo todo

	err := c.BindJSON(&updatedTodo)

	if err != nil {
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Task = updatedTodo.Task
			todos[i].Completed = updatedTodo.Completed
			c.IndentedJSON(http.StatusOK, todos[i])
			return
		}
	}
}

func deleteTodo(c *gin.Context) {
	id := c.Param("id")
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			c.IndentedJSON(http.StatusOK, todos)
			return
		}
	}
}
