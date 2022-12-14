package models

type Link struct {
	Model
	Code     string    `json:"string"`
	UserId   uint      `json:"user_id"`
	User     User      `json:"user" gorm:"foreignKey:UserId"`
	Products []Product `json:"product" gorm:"many2many:link_products"`
	Orders   []Order   `json:"orders" gorm:"-"`
}
