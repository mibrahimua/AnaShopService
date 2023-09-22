package controller

import (
	"AnaShopService/request"
	"AnaShopService/response"
	"AnaShopService/service"
	"github.com/gin-gonic/gin"
)

type ShopController struct {
	shopService *service.ShopService
}

func NewShopController(shopService *service.ShopService) *ShopController {
	return &ShopController{shopService}
}

// @Summary		Pay Cart Items
// @Description	Pay Cart Items
// @Produce		json
// @Param product_name body request.CartRequest false "product_name"
// @Success		200	{object} model.Product
// @Router			/cart [post]
func (uc *ShopController) CartRequest(c *gin.Context) {
	request := request.CartRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	err := uc.shopService.PaidCartItems(request.ID)
	if err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	c.JSON(200, response.Response{
		Status:  "success",
		Message: "success",
	})
}
