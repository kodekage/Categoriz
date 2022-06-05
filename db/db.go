package db

import (
	"errors"
	uuid "github.com/satori/go.uuid"
)

type Category struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Slug           string    `json:"slug"`
	ParentCategory *Category `json:"parentCategory"`
	isVisible      bool      `json:"isVisible"`
}

var perfumes = Category{
	Id:        uuid.NewV4(),
	Name:      "Perfumes",
	Slug:      "perfume",
	isVisible: true,
}

var categories = []Category{
	{
		Id:             uuid.FromStringOrNil("b2440ddc-910d-4ab5-b0dc-b01f3258eddc"),
		Name:           "Mild Perfumes",
		Slug:           "mild-perfumes",
		ParentCategory: &perfumes,
		isVisible:      true,
	},
	{
		Id:             uuid.FromStringOrNil("ed9b1c64-40e4-45a0-9a5f-8633b117e331"),
		Name:           "Strong Perfumes",
		Slug:           "strong-perfumes",
		ParentCategory: &perfumes,
		isVisible:      true,
	},
}

func FindAll() []Category {
	return categories
}

func Add(category Category) []Category {
	data := Category{
		Id:             uuid.NewV4(),
		Name:           category.Name,
		Slug:           category.Slug,
		isVisible:      category.isVisible,
		ParentCategory: category.ParentCategory,
	}

	categories = append(categories, data)

	return categories
}

func Find(id string) (Category, error) {
	var result Category
	var error error = nil

	for _, category := range categories {
		if id == category.Id.String() {
			result = category
		}
	}

	if (result == Category{}) {
		error = errors.New("Category not found")
	}

	return result, error
}
