package repository

import (
	"errors"
	// "wms_service/postgres"
	"wms_service/postgres"
	"wms_service/wms/responses"

	"github.com/gin-gonic/gin"
)



var DB *postgres.Db

func InitializeDB() {
	DB = postgres.GetCluster()
}

func CreateSku(c *gin.Context, sku *responses.CreateSkuCtrlResponse) error {
	if DB == nil {
		return errors.New("database connection is not initialized")
	}
	return DB.GetMasterDB(c).Create(sku).Error
}
func CreateHub(c *gin.Context,hub *responses.CreateHubCtrlResponse) error {
	if DB == nil {
		return errors.New("database connection is not initialized")
	}
	return DB.GetMasterDB(c).Create(hub).Error
}
func CreateInventory(c *gin.Context,inv *responses.CreateInventoryCtrlResponse) error {
	if DB == nil {
		return errors.New("database connection is not initialized")
	}
	return DB.GetMasterDB(c).Create(inv).Error
}