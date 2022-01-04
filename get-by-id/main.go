package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"dbhandler"
)

func getPeopleByID(c *gin.Context) {
	db := dbhandler.Init()
	id := c.Param("id")

	var person dbhandler.Person
	db.First(&person, id)

	if person == (dbhandler.Person{}) {
		c.IndentedJSON(http.StatusNotFound, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, person)
}

func main() {
	router := gin.Default()

	router.GET("/people/:id", getPeopleByID)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
