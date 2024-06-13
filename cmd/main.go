package main

import (
	"fmt"
	Location "forro_project/api/Location/Domain/Model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"os"
)

type Evento struct {
	Title     string
	Data      string
	Cidade    string
	Descricao string
}

func main() {
	startDbConnection()
	startServer()
}

func startDbConnection() *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable TimeZone=UTC",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Location.Country{}, &Location.State{}, &Location.City{}, &Location.Address{})

	return db
}

func startServer() {
	engine := gin.Default()
	engine = setupApp(engine)
	engine.Run(":8080")
}

func setupApp(engine *gin.Engine) *gin.Engine {
	t, err := template.ParseFiles("app/Views/Templates/Main/main.html")
	if err != nil {
		panic(err)
	}

	engine.SetHTMLTemplate(t)
	engine.StaticFS("/static", http.Dir("app/Views/Templates/Main"))

	return setupAppRouters(engine)
}

func setupAppRouters(engine *gin.Engine) *gin.Engine {
	appGroup := engine.Group("/app")
	appGroup.GET("", func(c *gin.Context) {
		eventos := []Evento{
			{Title: "Evento 1", Data: "2023-02-20", Cidade: "São Paulo", Descricao: "Lorem ipsum dolor sit ameadadasdt."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
			{Title: "Evento 2", Data: "2023-03-15", Cidade: "Rio de Janeiro", Descricao: "Lorem ipsum dolor sit amet."},
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

	return engine
}
