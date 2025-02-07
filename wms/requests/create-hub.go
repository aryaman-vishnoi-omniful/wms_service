package requests
type CreateHubCtrlRequest struct {
	HubId         string `json:"hub_id"`
	TenantID      string `json:"tenant_id"`
	Manager_email string `json:"manager_email"`
	ContactNo     string `json:"contactNo"`
	HubName       string `json:"hub_name"`
	Location      string `json:"location,omitempty"`
}
type GetHubsSvcRequest struct {
	TenantID string `json:"tenant_id"`
}