package api

import (
	"database/sql"
	"net/http"

	db "github.com/alekseiapa/apple_store/db/sqlc"
	"github.com/gin-gonic/gin"
)

// these api endpoints are inspired by https://developers.shopware.com/developers-guide/rest-api/examples/order/
type createOrderRequest struct {
	UserUuid    int64 `json:"user_uuid" binding:"required"`
	Quantity    int32 `json:"quantity" binding:"required"`
	ProductUuid int64 `json:"product_uuid" binding:"required"`
}

type orderResponse struct {
	Uuid     int64 `json:"order_uuid"`
	UserUuid int64 `json:"user_uuid"`
	Quantity int64 `json:"quantity"`
}

func newOrderResponse(order db.Order) orderResponse {
	return orderResponse{
		Uuid:     order.Uuid,
		UserUuid: order.UserUuid,
		Quantity: order.Quantity,
	}
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

type getOrderRequest struct {
	Uuid int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getOrder(ctx *gin.Context) {
	var req getOrderRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	order, err := server.store.GetOrder(ctx, req.Uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newOrderResponse(order)
	ctx.JSON(http.StatusOK, rsp)
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
		ctx.JSON(http.StatusNotFound, notFoundResponse("order"))
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
