package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"dbhandler"
)

func postPeople(c *gin.Context) {
	db := dbhandler.Init()

	var person dbhandler.Person

	if err := c.BindJSON(&person); err != nil {
		log.Fatalln(err)
		return
	}

	db.Create(&person)

	c.IndentedJSON(http.StatusCreated, person)
}

func main() {
	router := gin.Default()

	router.POST("/people", postPeople)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
