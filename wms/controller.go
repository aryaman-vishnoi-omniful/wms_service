package wms

import (
	"log"
	"net/http"
	"wms_service/wms/requests"
	// "wms_service/wms/responses"
	"wms_service/wms/services"

	"github.com/gin-gonic/gin"
	// commonError "github.com/omniful/go_commons/error"
	// 	oresponse "github.com/omniful/go_commons/response"
)

type WMSController struct {
	// services services.services
}

func (wc *WMSController) CreateSku(c *gin.Context) {
	var req *requests.CreateSkuCtrlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	

	svcResp, cusErr := services.CreateSku(c, req)
	if cusErr!=nil {
		// commonError.NewErrorResponse(c, cusErr)
		log.Fatal(cusErr.Error())
		return
	}
	c.JSON(http.StatusCreated,gin.H{"Sku created":svcResp})


}
func (wc *WMSController) CreateHub(c *gin.Context) {
	var req *requests.CreateHubCtrlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		
		return
	}


	svcResp, cusErr := services.CreateHub(c, req)
	if cusErr!=nil {
	
		log.Fatal(cusErr.Error())
		return
	}

	c.JSON(http.StatusCreated,gin.H{"hub created":svcResp})
}

func (wc *WMSController) CreateInventory(c *gin.Context) {
	var req *requests.CreateInventoryCtrlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}



	svcResp, cusErr := services.CreateInventory(c, req)
	if cusErr!=nil {
		log.Fatal(cusErr.Error())
		return
	}

	c.JSON(http.StatusCreated,gin.H{"inventory created":svcResp})
}

func (wc *WMSController) GetInventory(c *gin.Context) {

	sellerID := c.Query("seller_id")
	hubID := c.Query("hub_id")

	if sellerID == "" || hubID == "" {
		c.JSON(http.StatusBadRequest,gin.H{"error":"hub id or seller id required"})
		return
	}

	svcReq := &requests.GetInventorySvcRequest{
		SellerID: sellerID,
		HubID:    hubID,
	}

	svcResp, cusErr := services.GetInventory(c, svcReq)
	if cusErr!=nil {
		// commonError.NewErrorResponse(c, cusErr)
		log.Fatal(cusErr.Error())
		return
	}

	c.JSON(http.StatusCreated,gin.H{"The inventory for this hub :":svcResp})
}

func (wc *WMSController) DeductInventory(c *gin.Context) {
	var req *requests.AdjustInventoryCtrlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}


	svcResp, cusErr := services.DeductInventory(c, req)
	if cusErr!=nil {
		
		log.Fatal(cusErr.Error())
		return
	}

	c.JSON(http.StatusCreated,gin.H{"inventory updated":svcResp})
}

func (wc *WMSController) AddInventory(c *gin.Context) {
	var req *requests.AdjustInventoryCtrlRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}



	svcResp, cusErr := services.AddInventory(c, req)
	if cusErr!=nil {
		// commonError.NewErrorResponse(c, cusErr)
		log.Fatal(cusErr.Error())
		return
	}
	c.JSON(http.StatusCreated,gin.H{"inventory updated":svcResp})
}

func (wc *WMSController) GetHub(c *gin.Context) {
	hubID := c.Param("id")
	if hubID == "" {
		c.JSON(http.StatusBadRequest,gin.H{"error":"missing hub id"})
		return
	}

	svcResp, cusErr := services.GetHub(c, hubID)
	if cusErr!=nil{
		log.Fatal("could not get the hub")
		return
	}

	c.JSON(http.StatusOK,gin.H{"Details of the hub":svcResp})
}

func (wc *WMSController) GetHubs(c *gin.Context) {
	tenantID := c.Query("tenant_id") // optional filter
	svcReq := &requests.GetHubsSvcRequest{
		TenantID: tenantID,
	}

	svcResp, cusErr := services.GetHubs(c, svcReq)
	if cusErr!=nil{
		log.Fatal("could not get hubs")
		return

	}

	c.JSON(http.StatusOK,gin.H{"the hubs are ":svcResp})
}
func (wc *WMSController) GetSku(c *gin.Context) {
	// tenantID := c.Query("tenant_id")
	sellerID := c.Query("seller_id")
	skuCode := c.Query("sku_code")

	if sellerID == "" || skuCode == "" {
		c.JSON(http.StatusBadRequest,gin.H{"error":"missing fields"})
	}

	svcReq := &requests.GetSkuSvcRequest{
		SellerID: sellerID,
		SkuCode:  skuCode,
	}

	svcResp, cusErr := services.GetSku(c, svcReq)
	if cusErr!=nil{
		log.Fatal("could not get sku")
		return
	}

	c.JSON(http.StatusOK,gin.H{"the sku details are ":svcResp})
}
func (wc *WMSController) GetSkuById(c *gin.Context) {
	skuID := c.Param("id")
	if skuID == "" {
		c.JSON(http.StatusBadRequest,gin.H{"eror":"include id"})
	}

	

	svcResp, cusErr := services.GetSkuById(c, skuID)
	if cusErr!=nil{
		log.Fatal("could not get sku")
		return
	}

	c.JSON(http.StatusOK,gin.H{"fetched sku details":svcResp})
}



