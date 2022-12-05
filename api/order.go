package api

import (
	"database/sql"
	"net/http"

	db "github.com/alekseiapa/apple_store/db/sqlc"
	"github.com/gin-gonic/gin"
)

// these api endpoints are inspired by https://developers.shopware.com/developers-guide/rest-api/examples/order/
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

	if !server.validUser(ctx, req.UserUuid) {
		return
	}

	if !server.validProduct(ctx, req.ProductUuid) {
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

type deleteOrderRequest struct {
	Uuid int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteOrder(ctx *gin.Context) {
	var req deleteOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	r, err := server.store.DeleteOrder(ctx, req.Uuid)
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

func (server *Server) validUser(ctx *gin.Context, UserUuid int64) bool {
	_, err := server.store.GetUser(ctx, UserUuid)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	return true
}

func (server *Server) validProduct(ctx *gin.Context, ProductUuid int64) bool {
	_, err := server.store.GetProduct(ctx, ProductUuid)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	return true
}
