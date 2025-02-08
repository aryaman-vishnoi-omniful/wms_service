package responses

import "gorm.io/gorm"
type CreateSkuCtrlResponse struct {
	gorm.Model
	SellerID    string `gorm:"column:seller_id;type:varchar(100);not null" json:"seller_id"`
	SkuCode     string `gorm:"column:sku_code;type:varchar(100);uniqueIndex;not null" json:"sku_code"`
	Name        string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Price       int    `gorm:"column:price;not null" json:"price"`
	Dimensions  string `gorm:"column:dimensions;type:varchar(100)" json:"dimensions"`
	Fragile     string `gorm:"column:fragile;type:varchar(10)" json:"fragile"`
	Description string `gorm:"column:description;type:text" json:"description,omitempty"`
}
func (CreateSkuCtrlResponse) TableName() string {
	return "skus"
}



type GetSkuCtrlResponse struct {
	gorm.Model
	SellerID    string `json:"seller_id"`
	SkuCode     string `json:"sku_code"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Dimensions  string `json:"dimensions"`
	Fragile     string `json:"fragile"`
	Description string `json:"description,omitempty"`
}
func (GetSkuCtrlResponse) TableName() string {
    return "skus"
}
