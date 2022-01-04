package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"dbhandler"
)

func getPeople(c *gin.Context) {
	db := dbhandler.Init()
	name := c.Query("name")
	age := c.Query("age")

	var people []dbhandler.Person

	if name != "" {
		db = db.Where("name = ?", name)
	}

	if age != "" {
		db = db.Where("age = ?", age)
	}

	db.Order("id").Find(&people)

	if len(people) == 0 {
		c.IndentedJSON(http.StatusNotFound, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, people)
}

func main() {
	router := gin.Default()

	router.GET("/people", getPeople)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
