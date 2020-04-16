package main

import (
	"log"
	"net/http"
	"os"

	lifestyle "github.com/pennz/antlr_lifestyle"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	log.Println("Check before")
	log.Println("Preparation Done")
	log.Println(lifestyle.Actions)
	router.Run(":" + port)
	log.Fatal("Check, router is down")
}
