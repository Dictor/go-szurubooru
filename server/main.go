package main

import (
	"github.com/dictor/go-szurubooru/model"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	e := echo.New()
	e.GET("/tag", model.ListingModel(model.Tag{}))
	e.POST("/tag", model.CreateModel(model.Tag{}))
	log.Fatal(e.Start(":80"))
}
