package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Images      pq.StringArray `gorm:"type:text[]"`
}

func New(name, description string, images pq.StringArray) *Product {
	return &Product{
		Name:        name,
		Description: description,
		Images:      images,
	}
}
