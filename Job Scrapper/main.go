package main

import (
	"os"
	"strings"

	"github.com/labstack/echo"
)

const File_Name string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove("File_Name")
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)

	return c.Attachment("File_Name", "File_Name") // 첨부파일 return(jobs.csv파일을 사용자에게 jobs.csv로 전달)

}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))

}
