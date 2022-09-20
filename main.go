package main

import (
	"log"
	"pricing/database"
	"pricing/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Pricing server is starting...")

	d, err := database.New()
	if err != nil {
		panic(err)
	}
	h := handlers.New(d)

	r := gin.Default()

	r.POST("/createRule", h.CreateRule)
	r.GET("/changePrice", h.ChangePrice)

	log.Println("Listening and serving HTTP on 127.0.0.1:8080")

	r.Run("0.0.0.0:8080")
}
