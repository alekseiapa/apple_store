package api

import (
	"net/http"

	db "github.com/alekseiapa/apple_store/db/sqlc"
	"github.com/gin-gonic/gin"
)

// these api endpoints are inspired from https://developers.shopware.com/developers-guide/rest-api/examples/order/

type createOrderRequest struct {
	UserUuid    int64 `json:"UserUuid" binding:"required"`
	Quantity    int32 `json:"Quantity" binding:"required"`
	ProductUuid int64 `json:"ProductUuid" binding:"required"`
}

func (server *Server) createOrder(ctx *gin.Context) {
	var req createOrderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.BuyProductTxParams{
		UserUuid:    req.UserUuid,
		Quantity:    req.Quantity,
		ProductUuid: req.ProductUuid,
	}
	order, err := server.store.BuyProductTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, order)
}
