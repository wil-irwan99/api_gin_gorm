package controller

import (
	"api_gin_ref/delivery/api"
	"api_gin_ref/model"
	"api_gin_ref/usecase"
	"api_gin_ref/utils"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	router       *gin.Engine
	ucProduct    usecase.CreateProductUseCase
	ucGetProduct usecase.GetProductUseCase
	api.BaseApi
}

func (p *ProductController) createNewProduct(c *gin.Context) {
	var newProduct model.Product
	err := p.ParseRequestBody(c, &newProduct)
	if err != nil {
		p.Failed(c, utils.RequiredError())
		return
	}
	err = p.ucProduct.CreateProduct(&newProduct)
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, newProduct)

	// if err := c.BindJSON(&newProduct); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"status":  "BAD REQUEST",
	// 		"message": err.Error(),
	// 	})
	// } else {
	// 	err := p.ucProduct.CreateProduct(newProduct)
	// 	if err != nil {
	// 		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	// 			"status":  "FAILED",
	// 			"message": "Error when creating Product",
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status":  "SUCCESS",
	// 		"message": newProduct,
	// 	})
	// }
}

func (p *ProductController) getProductList(c *gin.Context) {
	products, err := p.ucGetProduct.GetProduct()
	if err != nil {
		p.Failed(c, err)
		return
	}
	p.Success(c, products)

	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
	// 		"status":  "FAILED",
	// 		"message": "Error when getting all Product",
	// 	})
	// 	return
	// }
	// c.JSON(http.StatusOK, gin.H{
	// 	"status":  "SUCCESS",
	// 	"message": res,
	// })
}

func NewProductController(router *gin.Engine, ucProduct usecase.CreateProductUseCase, ucGetProduct usecase.GetProductUseCase) *ProductController {
	//Di sini akan terdapat kumpulan semua request method yang dibutuhkan
	controller := ProductController{
		router:       router,
		ucProduct:    ucProduct,
		ucGetProduct: ucGetProduct,
	}

	//ini method-methodnya
	router.POST("/product", controller.createNewProduct)
	router.GET("/product", controller.getProductList)
	return &controller

}
