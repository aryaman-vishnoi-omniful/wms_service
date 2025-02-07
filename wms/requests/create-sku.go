package requests
type CreateSkuCtrlRequest struct {
	SellerID    string `json:"seller_id"`
	SkuCode     string `json:"sku_code"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Dimensions  string `json:"dimensions"`
	Fragile     string `json:"fragile"`
	Description string `json:"description,omitempty"`
}

type GetSkuSvcRequest struct {
	SellerID string `json:"seller_id"`
	SkuCode  string `json:"sku_code"`
}