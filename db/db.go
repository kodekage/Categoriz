package db

import (
	"errors"
	uuid "github.com/satori/go.uuid"
)

type Category struct {
	Id             uuid.UUID `json:"id,omitempty"`
	Name           string    `json:"name,omitempty" validate:"required,min=4"`
	Slug           string    `json:"slug,omitempty" validate:"required_with=Name"`
	ParentCategory uuid.UUID `json:"parentCategory,omitempty" validate:"required"`
	IsVisible      bool      `json:"isVisible,omitempty"`
}

var perfumes = Category{
	Id:        uuid.NewV4(),
	Name:      "Perfumes",
	Slug:      "perfumes",
	IsVisible: true,
}

var categories = []Category{
	{
		Id:             uuid.FromStringOrNil("b2440ddc-910d-4ab5-b0dc-b01f3258eddc"),
		Name:           "Mild Perfumes",
		Slug:           "mild-perfumes",
		ParentCategory: perfumes.Id,
		IsVisible:      true,
	},
	{
		Id:             uuid.FromStringOrNil("ed9b1c64-40e4-45a0-9a5f-8633b117e331"),
		Name:           "Strong Perfumes",
		Slug:           "strong-perfumes",
		ParentCategory: perfumes.Id,
		IsVisible:      true,
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
		IsVisible:      category.IsVisible,
		ParentCategory: category.ParentCategory,
	}

	categories = append(categories, data)

	return categories
}

func FindByIdOrSlug(identifier string) (error, Category) {
	var result Category
	var err error = nil

	for _, category := range categories {
		if identifier == category.Id.String() || identifier == category.Slug {
			result = category
		}
	}

	if (result == Category{}) {
		err = errors.New("category not found")
	}

	return err, result
}

func FindIndex(id string) (error, int) {
	var index int
	var result Category
	var err error = nil

	for i, category := range categories {
		if id == category.Id.String() {
			index = i
			result = category
		}
	}

	if (result == Category{}) {
		err = errors.New("category not found")
	}

	return err, index
}

func UpdateById(id string, category Category) (error, Category) {
	var err error = nil
	resultErr, index := FindIndex(id)

	if resultErr != nil {
		return resultErr, Category{}
	}

	categories[index].Name = category.Name
	categories[index].IsVisible = category.IsVisible
	categories[index].Slug = category.Slug
	categories[index].ParentCategory = category.ParentCategory

	return err, categories[index]
}

func Delete(id string) (error, []Category) {
	var error error
	err, index := FindIndex(id)

	if err != nil {
		error = err

		return error, nil
	}

	categories = append(categories[:index], categories[index+1:]...)

	return error, categories
}
