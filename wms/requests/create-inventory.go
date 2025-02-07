package requests
type CreateInventoryCtrlRequest struct {
	// InvID    string `json:"inv_id"`
	HubID    string `json:"hub_id"`
	SkuId    string `json:"sku_id"`
	SellerID string `json:"seller_id"`
	Quantity int    `json:"quantity"`
}

type AdjustInventoryCtrlRequest struct {
	TenantID string `json:"tenant_id" binding:"required"`
	SellerID string `json:"seller_id" binding:"required"`
	HubID    string `json:"hub_id" binding:"required"`
	SkuCode  string `json:"sku_code" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type GetInventorySvcRequest struct {
	SellerID string `json:"seller_id"`
	HubID    string `json:"hub_id"`
}
