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

func NewShopController(productService *service.ShopService) *ShopController {
	return &ShopController{productService}
}

// @Summary		Pay Cart Items
// @Description	Pay Cart Items
// @Produce		json
// @Param product_name body request.ProductRequest false "product_name"
// @Success		200	{object} model.Product
// @Router			/product [post]
func (uc *ShopController) GetListProductByName(c *gin.Context) {
	request := request.ProductRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	data, err := uc.shopService.GetListProductByName(request.Name)
	if err != nil {
		c.JSON(400, response.Response{
			Status:  "error",
			Message: err.Error(),
		})
	}

	c.JSON(200, response.Response{
		Status:  "success",
		Data:    data,
		Message: "success",
	})
}
