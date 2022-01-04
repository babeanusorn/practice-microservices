package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"dbhandler"
)

func patchPeople(c *gin.Context) {
	db := dbhandler.Init()
	id := c.Param("id")

	var person dbhandler.Person
	db.First(&person, id)

	if person == (dbhandler.Person{}) {
		c.IndentedJSON(http.StatusNotFound, nil)
		return
	}

	if err := c.BindJSON(&person); err != nil {
		log.Fatalln(err)
		return
	}

	db.Save(&person)

	c.IndentedJSON(http.StatusOK, person)
}

func main() {
	router := gin.Default()

	router.PATCH("/people/:id", patchPeople)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
