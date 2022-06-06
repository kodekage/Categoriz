package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/kodekage/categoriz/db"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

var v = validator.New()

func main() {
	server := echo.New()
	port := os.Getenv("APP_PORT")
	//v := validator.New()

	server.HTTPErrorHandler = customHttpErrorHandler

	server.GET("/categories", getCategories)
	server.GET("/categories/:id", getCategoryById)
	server.POST("/categories", createCategory)
	server.PUT("/categories/:id", updateCategory)

	server.Logger.Fatal(server.Start(":" + port))
}

func getCategories(context echo.Context) error {
	return context.JSON(http.StatusOK, db.FindAll())
}

func getCategoryById(context echo.Context) error {
	id := context.Param("id")
	err, category := db.Find(id)

	if err != nil {
		return newHTTPError(http.StatusNotFound, "NotFound", err.Error())
	}

	return context.JSON(http.StatusOK, category)
}

func createCategory(context echo.Context) error {
	var reqestBody db.Category

	if err := context.Bind(&reqestBody); err != nil {
		return newHTTPError(http.StatusBadRequest, "BadRequest", err.Error())
	}

	if err := v.Struct(reqestBody); err != nil {
		return newHTTPError(http.StatusBadRequest, "BadRequest", err.Error())
	}

	result := db.Add(reqestBody)

	return context.JSON(http.StatusOK, result)
}

func updateCategory(context echo.Context) error {
	id := context.Param("id")
	var reqestBody db.Category

	if err := context.Bind(&reqestBody); err != nil {
		return newHTTPError(http.StatusBadRequest, "BadRequest", err.Error())
	}

	if err := v.Struct(reqestBody); err != nil {
		return newHTTPError(http.StatusBadRequest, "BadRequest", err.Error())
	}

	result, err := db.UpdateById(id, reqestBody)

	if err != nil {
		return newHTTPError(http.StatusBadRequest, "BadRequest", err.Error())
	}

	return context.JSON(http.StatusOK, result)
}
