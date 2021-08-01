package main

import (
	"net/http"
	"os"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Person struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Title   string `json:"title"`
	Account string `json:"account"`
}

func main() {
	e := echo.New()
	gofakeit.Seed(0)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		gofakeit.Seed(0)
		name := gofakeit.Name()
		phone := gofakeit.Phone()
		email := gofakeit.Email()
		title := gofakeit.JobTitle()
		account := gofakeit.AchAccount()

		person := &Person{
			Name:    name,
			Phone:   phone,
			Email:   email,
			Title:   title,
			Account: account,
		}
		return c.JSONPretty(http.StatusOK, person, " ")
	})

	httpPort := os.Getenv("HTTP_PORT")
	e.Logger.Fatal(e.Start(":" + httpPort))
}
