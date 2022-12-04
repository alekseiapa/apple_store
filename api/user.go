package api

import (
	"database/sql"
	"log"
	"net/http"

	db "github.com/alekseiapa/apple_store/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	FirstName  string `json:"FirstName" binding:"required"`
	MiddleName string `json:"MiddleName" binding:"required"`
	LastName   string `json:"LastName" binding:"required"`
	Gender     string `json:"Gender" binding:"required,oneof=M F"`
	Age        int16  `json:"Age" binding:"required"`
	Currency   string `json:"Currency" binding:"required,oneof=USD EUR RUB"`
	Balance    int64  `json:"Balance" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// TODO: Balance will be stored as USD in DB but for the User will be displayed according to his needs
	arg := db.CreateUserParams{
		FirstName:  req.FirstName,
		MiddleName: req.MiddleName,
		LastName:   req.LastName,
		Gender:     req.Gender,
		Age:        req.Age,
		Balance:    req.Balance,
	}
	user, err := server.store.CreateUser(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type getUserRequest struct {
	Uuid int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5"`
}

func (server *Server) listUser(ctx *gin.Context) {
	var req listUserRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	users, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}

type updateUserRequestUri struct {
	Uuid int64 `uri:"id" binding:"required,min=1"`
}
type updateUserRequestJson struct {
	FirstName  string `json:"FirstName" binding:"required"`
	MiddleName string `json:"MiddleName" binding:"required"`
	LastName   string `json:"LastName" binding:"required"`
	Gender     string `json:"Gender" binding:"required,oneof=M F"`
	Age        int16  `json:"Age" binding:"required"`
	Currency   string `json:"Currency" binding:"required,oneof=USD EUR RUB"`
	Balance    int64  `json:"Balance" binding:"required"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var reqUri updateUserRequestUri
	var reqJson updateUserRequestJson

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&reqJson); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateUserParams{
		Uuid:       reqUri.Uuid,
		FirstName:  reqJson.FirstName,
		MiddleName: reqJson.MiddleName,
		LastName:   reqJson.LastName,
		Gender:     reqJson.Gender,
		Age:        reqJson.Age,
		Balance:    reqJson.Balance,
	}
	user, err := server.store.UpdateUser(ctx, arg)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type deleteUserRequest struct {
	Uuid int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteUser(ctx *gin.Context) {
	var req deleteUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	r, err := server.store.DeleteUser(ctx, req.Uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	log.Println(r)
	if r == 0 {
		ctx.JSON(http.StatusNotFound, notFoundResponse())
		return
	}
	ctx.JSON(http.StatusOK, successDeleteResponse())
}
