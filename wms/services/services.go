package services

import (
	"wms_service/wms/repository"
	"wms_service/wms/requests"
	"wms_service/wms/responses"

	"github.com/gin-gonic/gin"
)

func CreateSku(c *gin.Context, req *requests.CreateSkuCtrlRequest) (*responses.CreateSkuCtrlResponse, error) {

	// var skumodel responses.CreateSkuCtrlResponse
	sku := &responses.CreateSkuCtrlResponse{
		// SkuID: req.SkuCode,
		SellerID:    req.SellerID,
		SkuCode:     req.SkuCode,
		Name:        req.Name,
		Price:       req.Price,
		Fragile:     req.Fragile,
		Dimensions:  req.Dimensions,
		Description: req.Description,
	}
	err := repository.CreateSku(c,sku)
	if err != nil {
		return nil,err
	}
	return sku, nil

}

func CreateHub(c *gin.Context, req *requests.CreateHubCtrlRequest) (*responses.CreateHubCtrlResponse, error) {

	// var skumodel responses.CreateSkuCtrlResponse
	hub := &responses.CreateHubCtrlResponse{
		// SkuID: req.Sku,
		// HubId: req.HubId,
        HubName: req.HubName,
        ContactNo: req.ContactNo,
        Location: req.Location,
        TenantID: req.TenantID,
        Manager_email: req.Manager_email,


	}
	err := repository.CreateHub(c,hub)
	if err != nil {
		return nil,err
	}
	return hub, nil

}
func CreateInventory(c *gin.Context, req *requests.CreateInventoryCtrlRequest) (*responses.CreateInventoryCtrlResponse, error) {

	// var skumodel responses.CreateSkuCtrlResponse
	inv := &responses.CreateInventoryCtrlResponse{
        Id: req.Id,
        HubID: req.HubID,
        SkuId: req.SkuId,
        Quantity: req.Quantity,

	

	}
	err := repository.CreateInventory(c,inv)
	if err != nil {
		return nil,err
	}
	return inv, nil

}

