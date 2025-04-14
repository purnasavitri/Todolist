package model

type Category struct {
    ID          uint   `gorm:"column:id_category;primaryKey" json:"id"`
    CategoryName string `gorm:"column:category" json:"category"`
}
func (Category) TableName() string {
	return "category"
}