package meal

import "gorm.io/gorm"

type Meal struct {
	gorm.Model
	Name  string
	Price uint //price in cents
}
