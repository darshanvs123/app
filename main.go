package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type person struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Depart string `json:"depart"`
	Clg    string `json:"clg"`
}

var persons = []person{
	{Id: "0", Name: "vijay", Depart: "IT", Clg: "MEC"},
	{Id: "1", Name: "arjun", Depart: "IT", Clg: "MEC"},
	{Id: "2", Name: "vicky", Depart: "IT", Clg: "MEC"},
}

func getPerson(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, persons)
}

// func createPerson(c *gin.Context) {
// 	var newPerson person
// 	if err := c.BindJSON(&newPerson); err != nil {
// 		return
// 	}
// 	c.IndentedJSON(http.StatusCreated, newPerson)
// }

func main() {
	router := gin.Default()
	router.GET("persons/", getPerson)
	// router.POST("persons/", createPerson)
	// router.GET("persons/:id", PersonById)
	router.Run("localhost:8080")
}
