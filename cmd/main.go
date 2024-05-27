package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Evento struct {
	Title     string
	Data      string
	Cidade    string
	Descricao string
}

func main() {
	r := gin.Default()

	// Load HTML templates from app/Views/Templates/Main
	t, err := template.ParseFiles("app/Views/Templates/Main/main.html")
	if err != nil {
		panic(err)
	}

	r.SetHTMLTemplate(t)

	// Serve static files from app/Views/Templates/Main
	r.StaticFS("/static", http.Dir("app/Views/Templates/Main"))

	r.GET("/", func(c *gin.Context) {
		eventos := []Evento{
			{Title: "Evento 1", Data: "2023-02-20", Cidade: "São Paulo", Descricao: "Lorem ipsum dolor sit ameadadasdt."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			// Add more eventos here
		}

		c.HTML(http.StatusOK, "main.html", gin.H{
			"eventos": eventos,
		})

		// Set cache control headers
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
	})

	r.Run(":8080")
}
