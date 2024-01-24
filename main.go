package main

import (
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID          string `json:"id"`
	Item        string `json:"item"`
	IsCompleted bool   `json:"isCompleted"`
}

var todos = []todo{
{ID: "1", Item: "Clean Room", IsCompleted: false},
{ID: "2", Item: "Clean Kitchen", IsCompleted: false},
{ID: "3", Item: "Clean Toilet", IsCompleted: false},
}

func getAllTodos(context *gin.Context)  {
	context.JSON(http.StatusOK, todos)
	return 
}

func createTodo(context *gin.Context)  {
	var newTodo todo; // the variable to hold request body.

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)
	context.JSON(http.StatusCreated,newTodo)
	return
}

func getOneTodo(context *gin.Context)  { 
	id:= context.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return 
	}

	context.JSON(http.StatusOK, todo)
	return
}

func getTodoById(id string) (*todo, error) {
	for i, t:= range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("Todo not found")
}
func main()  {
	router:= gin.Default();
	router.GET("/todos", getAllTodos )
	router.GET("/todos/:id", getOneTodo )
	router.POST("/todos", createTodo )
	router.Run("localhost:4000") 
}