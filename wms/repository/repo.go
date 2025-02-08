package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"wms_service/postgres"
	"wms_service/redis"
	"wms_service/wms/requests"
	"wms_service/wms/responses"
)

var DB *postgres.Db

func InitializeDB() {
	DB = postgres.GetCluster()
}
func CreateSku(ctx context.Context, sku *responses.CreateSkuCtrlResponse) error {
	if DB == nil {
		return errors.New("database connection is not initialized")
	}
	return DB.GetMasterDB(ctx).Create(sku).Error
}

func CreateHub(ctx context.Context, hub *responses.CreateHubCtrlResponse) error {
	if DB == nil {
		return errors.New("database connection is not initialized")
	}
	return DB.GetMasterDB(ctx).Create(hub).Error
}

func CreateInventory(ctx context.Context, inv *responses.CreateInventoryCtrlResponse) error {
	if DB == nil {
		return errors.New("database connection is not initialized")
	}
	return DB.GetMasterDB(ctx).Create(inv).Error
}

func GetInventory(ctx context.Context, req *requests.GetInventorySvcRequest) (*responses.GetInventoryCtrlResponse, error) {
	if DB == nil {
		return nil, errors.New("database connection is not initialized")
	}
	var inventories []responses.CreateInventoryCtrlResponse
	err := DB.GetMasterDB(ctx).Where("seller_id = ? AND hub_id = ?", req.SellerID, req.HubID).Find(&inventories).Error
	if err != nil {
		return nil, err
	}
	items := []responses.InventoryItem{}
	for _, inv := range inventories {
		items = append(items, responses.InventoryItem{
			SkuCode:   inv.SkuId, 
			HubID:     inv.HubID,
			Inventory: inv.Quantity,
		})
	}
	return &responses.GetInventoryCtrlResponse{InventoryItems: items}, nil
}

func DeductInventory(ctx context.Context, req *requests.AdjustInventoryCtrlRequest) (*responses.AdjustInventoryCtrlResponse, error) {
	if DB == nil {
		return nil, errors.New("database connection is not initialized")
	}
	err := DB.GetMasterDB(ctx).Exec(
		"UPDATE inventories SET quantity = quantity - ? WHERE hub_id = ? AND sku_id = ? AND seller_id = ?",
		req.Quantity, req.HubID, req.SkuCode, req.SellerID,
	).Error
	if err != nil {
		return nil, err
	}
	return &responses.AdjustInventoryCtrlResponse{
		HubID:    req.HubID,
		SkuId:    req.SkuCode,
		Quantity: req.Quantity,
	}, nil
}

func AddInventory(ctx context.Context, req *requests.AdjustInventoryCtrlRequest) (*responses.AdjustInventoryCtrlResponse, error) {
	if DB == nil {
		return nil, errors.New("database connection is not initialized")
	}
	err := DB.GetMasterDB(ctx).Exec(
		"UPDATE inventories SET quantity = quantity + ? WHERE hub_id = ? AND sku_id = ? AND seller_id = ?",
		req.Quantity, req.HubID, req.SkuCode, req.SellerID,
	).Error
	if err != nil {
		return nil, err
	}
	return &responses.AdjustInventoryCtrlResponse{
		HubID:    req.HubID,
		SkuId:    req.SkuCode,
		Quantity: req.Quantity,
	}, nil
}
func GetHub(ctx context.Context, hubID string) (*responses.GetHubCtrlResponse, error) {
	if DB == nil {
		return nil, errors.New("database connection is not initialized")
	}
	var hub responses.GetHubCtrlResponse
	err := DB.GetMasterDB(ctx).Where("hub_id = ?", hubID).First(&hub).Error
	if err != nil {
		return nil, err
	}
	return &hub, nil
}



func GetHubs(ctx context.Context, req *requests.GetHubsSvcRequest) (*responses.GetHubsCtrlResponse, error) {
	if DB == nil {
		return nil, errors.New("database connection is not initialized")
	}
	var hubs []responses.GetHubCtrlResponse
	query := DB.GetMasterDB(ctx).Model(&responses.CreateHubCtrlResponse{})
	if req.TenantID != "" {
		query = query.Where("tenant_id = ?", req.TenantID)
	}
	err := query.Find(&hubs).Error
	if err != nil {
		return nil, err
	}
	return &responses.GetHubsCtrlResponse{Hubs: hubs}, nil
}

// Get SKU by query parameters
func GetSku(ctx context.Context, req *requests.GetSkuSvcRequest) (*responses.GetSkuCtrlResponse, error) {
	if DB == nil {
		return nil, errors.New("database connection is not initialized")
	}
	var sku responses.GetSkuCtrlResponse
	err := DB.GetMasterDB(ctx).Where("seller_id = ? AND sku_code = ?", req.SellerID, req.SkuCode).First(&sku).Error
	if err != nil {
		return nil, err
	}
	return &sku, nil
}

func GetSkuById(ctx context.Context, skuID string) (*responses.GetSkuCtrlResponse, error) {
	if DB == nil {
		return nil, errors.New("database connection is not initialized")
	}
	cacheKey := fmt.Sprintf("sku:%s", skuID)
	RD:=redis.GetClient()
	cached_data,err:=RD.Get(ctx,cacheKey)
	if err == nil {
		// Cache Hit: Unmarshal the cached JSON response
		var sku responses.GetSkuCtrlResponse
		if err := json.Unmarshal([]byte(cached_data), &sku); err == nil {
			fmt.Println()
			fmt.Println()
			fmt.Println("Cache hit for SKU:", skuID)
			fmt.Println()
			fmt.Println()
			return &sku, nil
		}
	}


	var sku responses.GetSkuCtrlResponse
	err = DB.GetMasterDB(ctx).Where("id = ?", skuID).First(&sku).Error
	if err != nil {
		return nil, err
	}
	jsonData, _ := json.Marshal(sku)
	RD.Set(ctx,cacheKey,string(jsonData),1*time.Minute)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("Cache miss, fetched from DB:", skuID)
	fmt.Println()
	fmt.Println()
	return &sku, nil
}
