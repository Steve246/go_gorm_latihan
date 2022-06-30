package repo

import (
	"go_gorms_lat/model"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error

	Update(model *model.Customer, data map[string]interface{}) error

	Delete(model *model.Customer) error
}

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) Delete(model *model.Customer) error {
	result := c.db.Delete(model).Error

	return result
}

func (c *customerRepository) Update(model *model.Customer, data map[string]interface{}) error {

	result := c.db.Model(model).Updates(data).Error

	return result

	// db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})

}

func (c *customerRepository) Create(customer *model.Customer) error {
	result := c.db.Create(customer).Error
	return result
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	repo := new(customerRepository)
	repo.db = db
	return repo
}
