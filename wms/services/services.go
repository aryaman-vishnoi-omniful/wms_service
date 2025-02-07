package services

import (
	"context"
	"wms_service/wms/repository"
	"wms_service/wms/requests"
	"wms_service/wms/responses"
)

func CreateSku(ctx context.Context, req *requests.CreateSkuCtrlRequest) (*responses.CreateSkuCtrlResponse, error) {
	sku := &responses.CreateSkuCtrlResponse{
		SellerID:    req.SellerID,
		SkuCode:     req.SkuCode,
		Name:        req.Name,
		Price:       req.Price,
		Dimensions:  req.Dimensions,
		Fragile:     req.Fragile,
		Description: req.Description,
	}
	err := repository.CreateSku(ctx, sku)
	if err != nil {
		return nil, err
	}
	return sku, nil
}

func CreateHub(ctx context.Context, req *requests.CreateHubCtrlRequest) (*responses.CreateHubCtrlResponse, error) {
	hub := &responses.CreateHubCtrlResponse{
		HubId:         req.HubId,
		TenantID:      req.TenantID,
		Manager_email: req.Manager_email,
		ContactNo:     req.ContactNo,
		HubName:       req.HubName,
		Location:      req.Location,
	}
	err := repository.CreateHub(ctx, hub)
	if err != nil {
		return nil, err
	}
	return hub, nil
}

func CreateInventory(ctx context.Context, req *requests.CreateInventoryCtrlRequest) (*responses.CreateInventoryCtrlResponse, error) {
	inv := &responses.CreateInventoryCtrlResponse{
		// Id:       req.InvID,
		HubID:    req.HubID,
		SkuId:    req.SkuId,
		SellerID: req.SellerID,
		Quantity: req.Quantity,
	}
	err := repository.CreateInventory(ctx, inv)
	if err != nil {
		return nil, err
	}
	return inv, nil
}

func GetInventory(ctx context.Context, req *requests.GetInventorySvcRequest) (*responses.GetInventoryCtrlResponse, error) {
	return repository.GetInventory(ctx, req)
}

func DeductInventory(ctx context.Context, req *requests.AdjustInventoryCtrlRequest) (*responses.AdjustInventoryCtrlResponse, error) {
	return repository.DeductInventory(ctx, req)
}

func AddInventory(ctx context.Context, req *requests.AdjustInventoryCtrlRequest) (*responses.AdjustInventoryCtrlResponse, error) {
	return repository.AddInventory(ctx, req)
}
func GetHub(ctx context.Context, hubID string) (*responses.GetHubCtrlResponse, error) {
	return repository.GetHub(ctx, hubID)
}

func GetHubs(ctx context.Context, req *requests.GetHubsSvcRequest) (*responses.GetHubsCtrlResponse, error) {
	return repository.GetHubs(ctx, req)
}

func GetSku(ctx context.Context, req *requests.GetSkuSvcRequest) (*responses.GetSkuCtrlResponse, error) {
	return repository.GetSku(ctx, req)
}
func GetSkuById(ctx context.Context, skuID string) (*responses.GetSkuCtrlResponse, error) {
	return repository.GetSkuById(ctx, skuID)
}
