package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID     int
	Name   string
	Email  string
	Role   string
	Status string
}

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
		users := []User{
			{ID: 1, Name: "Symbat", Email: "sinderella.0907@gmail.com", Role: "Admin", Status: "Active"},
		}

		c.HTML(http.StatusOK, "users.html", gin.H{
			"Title": "Пользователи",
			"Users": users,
		})
	})

	r.Run(":8080")
}
