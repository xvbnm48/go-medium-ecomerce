package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xvbnm48/go-medium-ecomerce/model"
	"github.com/xvbnm48/go-medium-ecomerce/repository"
)

type ProductHandler interface {
	GetProduct(*gin.Context)
	GetAllProduct(*gin.Context)
	AddProduct(*gin.Context)
	UpdateProduct(*gin.Context)
	DeleteProduct(*gin.Context)
}

type productHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler() ProductHandler {
	return &productHandler{
		repo: repository.NewProductRepository(),
	}
}

func (h *productHandler) GetAllProduct(ctx *gin.Context) {
	product, err := h.repo.GetAllProduct()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (h *productHandler) GetProduct(ctx *gin.Context) {
	prodStr := ctx.Param("product")
	prodID, err := strconv.Atoi(prodStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "Invalid product ID",
		})
		return
	}
	product, err := h.repo.GetProduct(prodID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (h *productHandler) AddProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	product, err := h.repo.AddProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (h *productHandler) UpdateProduct(ctx *gin.Context) {
	var product model.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	prodStr := ctx.Param("product")
	prodID, err := strconv.Atoi(prodStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": "Invalid product ID",
		})
		return
	}
	product.ID = uint(prodID)
	product, err = h.repo.UpdateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

func (h *productHandler) DeleteProduct(ctx *gin.Context) {
	var product model.Product
	prodStr := ctx.Param("product")
	prodID, err := strconv.Atoi(prodStr)
	product.ID = uint(prodID)
	product, err = h.repo.DeleteProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, product)
}
