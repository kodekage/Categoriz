package main

import (
	"github.com/kodekage/categoriz/db"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

func main() {
	server := echo.New()
	port := os.Getenv("APP_PORT")

	//server.HTTPErrorHandler = customHttpErrorHandler

	server.GET("/categories", getCategories)
	server.GET("/categories/:id", getCategoryById)
	server.POST("/categories", createCategory)

	server.Logger.Fatal(server.Start(":" + port))
}

func getCategories(context echo.Context) error {
	return context.JSON(http.StatusOK, db.FindAll())
}

func getCategoryById(context echo.Context) error {
	id := context.Param("id")
	category, err := db.Find(id)

	if err != nil {
		return newHTTPError(http.StatusNotFound, "NotFound", err.Error())
	}

	return context.JSON(http.StatusOK, category)
}

func createCategory(context echo.Context) error {
	var reqestBody db.Category

	if err := context.Bind(&reqestBody); err != nil {
		return err
	}

	result := db.Add(reqestBody)

	return context.JSON(http.StatusOK, result)
}
