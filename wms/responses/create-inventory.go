package responses

import "gorm.io/gorm"

type CreateInventoryCtrlResponse struct {
	gorm.Model
	InvId       string `gorm:"column:inv_id;type:varchar(100);uniqueIndex;not null" json:"inv_id"`
	HubID    string `gorm:"column:hub_id;type:varchar(100);not null" json:"hub_id"`
	SkuId    string `gorm:"column:sku_id;type:varchar(100);not null" json:"sku_id"`
	SellerID string `gorm:"column:seller_id;type:varchar(100);not null" json:"seller_id"`
	Quantity int    `gorm:"column:quantity;not null" json:"quantity"`
}
func (CreateInventoryCtrlResponse) TableName() string {
	return "inventories"
}

type AdjustInventoryCtrlResponse struct {
	HubID    string `json:"hub_id"`
	SkuId    string `json:"sku_id"`
	Quantity int    `json:"quantity"`
}

type GetInventoryCtrlResponse struct {
	InventoryItems []InventoryItem `json:"inventory_items"`
}

type InventoryItem struct {
	SkuCode   string `json:"sku_code"`
	HubID     string `json:"hub_id"`
	Inventory int    `json:"inventory"`
}