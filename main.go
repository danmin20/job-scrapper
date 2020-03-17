package main

import (
	"fmt"
	"github.com/danmin20/job-scrapper/scrapper"
	"github.com/labstack/echo"
	"strings"
)

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	return nil
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
	scrapper.Scrape("term")
}
