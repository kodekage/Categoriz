package main

import uuid "github.com/satori/go.uuid"

type Category struct {
	Id             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Slug           string    `json:"slug"`
	ParentCategory *Category `json:"parentCategory"`
	isVisible      bool      `json:"isVisible"`
}

func DB() []Category {
	perfumes := Category{
		Id:        uuid.NewV4(),
		Name:      "Perfumes",
		Slug:      "perfume",
		isVisible: true,
	}

	categories := []Category{
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

	return categories
}
