package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"dbhandler"
)

func deletePeople(c *gin.Context) {
	db := dbhandler.Init()
	id := c.Param("id")

	var person dbhandler.Person
	db.First(&person, id)

	db.Delete(&dbhandler.Person{}, id)

	if person == (dbhandler.Person{}) {
		c.IndentedJSON(http.StatusNoContent, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, person)
}

func main() {
	router := gin.Default()

	router.DELETE("/people/:id", deletePeople)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
