package model

type Customer struct {
	Id        string `gorm:"primaryKey"`
	Name      string `gorm:"size:50; not null"`
	Phone     string
	Balance   int
	Is_member int       `gorm:"default:1"`
	BaseModel BaseModel `gorm:"embedded"`
}
