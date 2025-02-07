package router

import (
	"context"
	"wms_service/wms"
	// "oms_service/orders"

	// "oms_service/orders/requests"

	// "oms_service/database"

	// "oms_service/redis"

	// "net/http"
	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/log"
	// "github.com/omniful/go_commons/jwt/private"
)

func Initialize(ctx context.Context, s *http.Server) (err error) {
	s.Engine.Use(log.RequestLogMiddleware(log.MiddlewareOptions{
		Format:      config.GetString(ctx, "log.format"),
		Level:       config.GetString(ctx, "log.level"),
		LogRequest:  config.GetBool(ctx, "log.request"),
		LogResponse: config.GetBool(ctx, "log.response"),
	}))

	wms_v1 := s.Engine.Group("/wms/v1")
	var WMSController *wms.WMSController
	
	wms_v1.POST("/create-sku", WMSController.CreateSku)
	wms_v1.POST("/create-hub",WMSController.CreateHub)
	wms_v1.POST("/create-inventory",WMSController.CreateInventory)
	// wms_v1.GET("/GetInventory",WMSController.GetInventory)
	// wms_v1.PUT("/inventory/deduct",WMSController.DeductInventory)
	// wms_v1.PUT("/inventory/add",WMSController.AddInventory)
	// wms_v1.GET("/getHub/:id",WMSController.GetHub)
	// wms_v1.GET("/getHubs",WMSController.GetHubs)
	// // wms_v1.GET("/getSkus",WMSController.GetSkus)
	// wms_v1.GET("/getSku",WMSController.GetSku)
	// wms_v1.GET("/getSkuById/:id",WMSController.GetSkuById)
	



	return nil

}
