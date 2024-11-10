package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

var todos = []todo{
	{ID: "1", Title: "first", Desc: "first todo"},
	{ID: "2", Title: "second", Desc: "second todo"},
	{ID: "3", Title: "third", Desc: "third todo"},
}

func getTodos(c *gin.Context){
	c.IndentedJSON(http.StatusOK,todos)
}

func getTodo(c *gin.Context){
	id := c.Param("id")

	for _,val := range todos {
		if val.ID == id {
			c.IndentedJSON(http.StatusOK, val)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "value not found!"})
}

func updateTodo(c *gin.Context){
	id := c.Param("id")

	var updatedTodo todo
	error := c.BindJSON(&updatedTodo)
	
	if error != nil {
		c.IndentedJSON(400, gin.H{"message" : "Invalid request!"})
		return
	}
	
	for i,val := range todos {
		if val.ID == id {
			todos[i].Title = updatedTodo.Title
			todos[i].Desc = 	updatedTodo.Desc
			c.IndentedJSON(200, todos[i])
			return
		}
	}

	c.JSON(404, gin.H{"message" : "value not found!"})
}

func postTodo(c *gin.Context){
	var newTodo todo
	error := c.BindJSON(&newTodo)

	if error!=nil {
		c.IndentedJSON(400, gin.H{"message" : "Invalid request!"})
		return
	}

	todos = append(todos,newTodo)

	c.IndentedJSON(201, newTodo)
}

func RemoveIndex(s []todo, index int) []todo {
	return append(s[:index], s[index+1:]...)
}

func deleteTodo(c *gin.Context){
	id := c.Param("id")

	for i,val := range todos {
		if val.ID == id {
			todos = RemoveIndex(todos, i)
			c.IndentedJSON(200, gin.H{"message" : "value deleted!"})
			return
		}
	}
}

func main() {

	r := gin.Default()

	r.GET("/", func(c *gin.Context){
		c.IndentedJSON(200, gin.H{
			"message" : "This is the root endpoint",
		})
	})

	r.GET("/todos", getTodos)
	r.GET("/todos/:id", getTodo)

	r.POST("/todos/addTodo/", postTodo)
	r.POST("/todos/updateTodo/:id", updateTodo)

	r.DELETE("/todos/deleteTodo/:id", deleteTodo)

	r.Run()
}