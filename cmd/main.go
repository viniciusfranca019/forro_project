package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PageData struct {
	Title   string
	Content string
	Items   []string
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./app/Views/Templates/**/*.html")

	router.GET("/", func(c *gin.Context) {
		data := PageData{
			Title:   "Example Page",
			Content: "Welcome to the example page!",
			Items:   []string{"Item 1", "Item 2", "Item 3"},
		}
		c.HTML(http.StatusOK, "List.html", gin.H{"data": data})
	})

	router.Run(":8080")
}
