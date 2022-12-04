package api

import (
	"database/sql"
	"log"
	"net/http"

	db "github.com/alekseiapa/apple_store/db/sqlc"
	"github.com/alekseiapa/apple_store/util"
	"github.com/gin-gonic/gin"
)

type productResponse struct {
	Price       float32
	Currency    string
	InStock     int32
	Description string
	Uuid        int64
}

type createProductRequest struct {
	Description string  `json:"Description" binding:"required"`
	Price       float32 `json:"Price" binding:"required"`
	InStock     int32   `json:"InStock" binding:"required"`
	Currency    string  `json:"Currency" binding:"required,oneof=USD EUR RUB"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	convPrice := util.ConvertCur(req.Currency, "USD", req.Price)
	arg := db.CreateProductParams{
		Description: req.Description,
		Price:       convPrice,
		InStock:     req.InStock,
	}
	product, err := server.store.CreateProduct(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	prodRespJson := productResponse{
		Description: product.Description,
		InStock:     product.InStock,
		Uuid:        product.Uuid,
		Currency:    req.Currency,
		Price:       req.Price,
	}
	ctx.JSON(http.StatusCreated, prodRespJson)
}

type getProductRequestUri struct {
	Uuid int64 `uri:"id" binding:"required,min=1"`
}
type getProductRequestQuery struct {
	Currency string `form:"currency" binding:"required,oneof=USD EUR RUB"`
}

func (server *Server) getProduct(ctx *gin.Context) {
	var reqUri getProductRequestUri
	var reqQuery getProductRequestQuery

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindQuery(&reqQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := server.store.GetProduct(ctx, reqUri.Uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	prodRespJson := productResponse{
		Description: product.Description,
		InStock:     product.InStock,
		Uuid:        product.Uuid,
		Currency:    reqQuery.Currency,
		Price:       util.ConvertCur("USD", reqQuery.Currency, product.Price),
	}
	ctx.JSON(http.StatusOK, prodRespJson)
}

type listProductRequest struct {
	PageID   int32  `form:"page_id" binding:"required,min=1"`
	PageSize int32  `form:"page_size" binding:"required,min=5"`
	Currency string `form:"currency" binding:"required,oneof=USD EUR RUB"`
}

func (server *Server) listProduct(ctx *gin.Context) {
	var req listProductRequest
	var respProducts []productResponse
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListProductsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	products, err := server.store.ListProducts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	for _, product := range products {
		respProducts = append(respProducts, productResponse{
			Description: product.Description,
			InStock:     product.InStock,
			Uuid:        product.Uuid,
			Currency:    req.Currency,
			Price:       util.ConvertCur("USD", req.Currency, product.Price),
		})
	}

	ctx.JSON(http.StatusOK, respProducts)
}

type updateProductRequestUri struct {
	Uuid int64 `uri:"id" binding:"required,min=1"`
}
type updateProductRequestJson struct {
	Description string  `json:"Description" binding:"required"`
	Price       float32 `json:"Price" binding:"required"`
	InStock     int32   `json:"InStock" binding:"required"`
	Currency    string  `json:"Currency" binding:"required,oneof=USD EUR RUB"`
}

func (server *Server) updateProduct(ctx *gin.Context) {
	var reqUri updateProductRequestUri
	var reqJson updateProductRequestJson

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&reqJson); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateProductParams{
		Uuid:        reqUri.Uuid,
		Description: reqJson.Description,
		Price:       util.ConvertCur(reqJson.Currency, "USD", reqJson.Price),
		InStock:     reqJson.InStock,
	}
	product, err := server.store.UpdateProduct(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	log.Println(product.InStock)
	prodRespJson := productResponse{
		Description: product.Description,
		InStock:     product.InStock,
		Uuid:        product.Uuid,
		Currency:    reqJson.Currency,
		Price:       reqJson.Price,
	}
	ctx.JSON(http.StatusOK, prodRespJson)
}

type deleteProductRequest struct {
	Uuid int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteProduct(ctx *gin.Context) {
	var req deleteProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	r, err := server.store.DeleteProduct(ctx, req.Uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if r == 0 {
		ctx.JSON(http.StatusNotFound, notFoundResponse())
		return
	}
	ctx.JSON(http.StatusOK, successDeleteResponse())

}
