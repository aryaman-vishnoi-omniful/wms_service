package responses

import "gorm.io/gorm"

type GetHubCtrlResponse struct {
	HubId string `json:"hub_id"`
	TenantID      string `json:"tenant_id"`
    Manager_email string `json:"manager_email"`
    ContactNo     string `json:"contactNo"`
    HubName       string `json:"hub_name"`
    Location      string `json:"location,omitempty"`
}

type GetHubsCtrlResponse struct {
	Hubs []GetHubCtrlResponse `json:"hubs"`
}

type CreateHubCtrlResponse struct {
	gorm.Model
	HubId         string `gorm:"column:hub_id;type:varchar(100);uniqueIndex;not null" json:"hub_id"`
	TenantID      string `gorm:"column:tenant_id;type:varchar(100);not null" json:"tenant_id"`
	Manager_email string `gorm:"column:manager_email;type:varchar(255);not null" json:"manager_email"`
	ContactNo     string `gorm:"column:contact_no;type:varchar(50);not null" json:"contactNo"`
	HubName       string `gorm:"column:hub_name;type:varchar(255);not null" json:"hub_name"`
	Location      string `gorm:"column:location;type:varchar(255)" json:"location,omitempty"`
}