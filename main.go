package main

import (
	"github.com/labstack/echo"
	"net/http"
	"os"
)

func main() {
	server := echo.New()
	port := os.Getenv("APP_PORT")

	server.GET("/categories", getCategories)
	server.GET("/categories/:id", getCategoryById)
	server.POST("/categories", createCategory)

	server.Logger.Fatal(server.Start(":" + port))
}

func getCategories(context echo.Context) error {
	categories := DB()

	return context.JSON(http.StatusOK, categories)
}

func getCategoryById(context echo.Context) error {
	// TODO: migrate to SQL DB
	categories := DB()
	id := context.Param("id")
	var result Category

	for _, category := range categories {
		if category.Id.String() == id {
			result = category
		}
	}

	if (result == Category{}) {
		return context.JSON(http.StatusBadRequest, "Category not found")
	}

	return context.JSON(http.StatusOK, result)
}

func createCategory(context echo.Context) error {
	db := DB()

	var reqestBody Category

	if err := context.Bind(&reqestBody); err != nil {
		return err
	}

	db = append(db, reqestBody)

	return context.JSON(http.StatusOK, db)
}
