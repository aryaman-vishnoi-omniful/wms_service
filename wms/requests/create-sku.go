package requests
type CreateSkuCtrlRequest struct {
	// TenantID    string `json:"tenant_id"`
	SellerID    string `json:"seller_id"`
	SkuCode     string `json:"sku_code"`
	Name        string `json:"name"`
	Price 		int `json:"price"`
	Dimensions 	string `json:"dimensions"`
	Fragile		string `json:"fragile"`
	Description string `json:"description,omitempty"`
}

type CreateSkuSvcRequest struct {
	// TenantID    string
	SellerID    string
	SkuCode     string
	Name        string
	Description string
}
type GetSkuSvcRequest struct {
	TenantID string
	SellerID string
	SkuCode  string
}