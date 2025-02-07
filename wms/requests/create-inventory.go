package requests
type CreateInventoryCtrlRequest struct {
	InvID    string `json:"inv_id"`
	HubID    string `json:"hub_id"`
	SkuId    string `json:"sku_id"`
	SellerID string `json:"seller_id"`
	Quantity int    `json:"quantity"`
}

type AdjustInventoryCtrlRequest struct {
	// TenantID string `json:"tenant_id" binding:"required"`
	SellerID string `json:"seller_id" `
	HubID    string `json:"hub_id" `
	SkuCode  string `json:"sku_code"`
	Quantity int    `json:"quantity"`
}

type GetInventorySvcRequest struct {
	SellerID string `json:"seller_id"`
	HubID    string `json:"hub_id"`
}
