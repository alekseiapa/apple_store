package api

import (
	"database/sql"
	"net/http"

	db "github.com/alekseiapa/apple_store/db/sqlc"
	"github.com/alekseiapa/apple_store/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createUserRequest struct {
	FirstName  string  `json:"first_name" binding:"required"`
	MiddleName string  `json:"middle_name" binding:"required"`
	LastName   string  `json:"last_name" binding:"required"`
	Gender     string  `json:"gender" binding:"required,oneof=M F"`
	Age        int16   `json:"age" binding:"required"`
	Balance    float32 `json:"balance" binding:"required"`
	Username   string  `json:"username" binding:"required,alphanum"`
	Password   string  `json:"password" binding:"required,min=6"`
}

type userResponse struct {
	Uuid       int64   `json:"uuid"`
	FirstName  string  `json:"first_name"`
	MiddleName string  `json:"middle_name"`
	LastName   string  `json:"last_name"`
	Gender     string  `json:"gender"`
	Age        int16   `json:"age"`
	Balance    float32 `json:"balance"`
	Username   string  `json:"username"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Uuid:       user.Uuid,
		FirstName:  user.FirstName,
		MiddleName: user.MiddleName,
		LastName:   user.LastName,
		Gender:     user.Gender,
		Age:        user.Age,
		Balance:    user.Balance,
		Username:   user.Username,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		FirstName:      req.FirstName,
		MiddleName:     req.MiddleName,
		LastName:       req.LastName,
		Gender:         req.Gender,
		Age:            req.Age,
		Balance:        req.Balance,
		Username:       req.Username,
		HashedPassword: hashedPassword,
	}
	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, userNameExistsResponse())
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newUserResponse(user)
	ctx.JSON(http.StatusCreated, rsp)
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
	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// if user.Username != authPayload.Username {
	// 	err := errors.New("can't buy on a behalf of other user")
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type listUserRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5"`
}

func (server *Server) listUser(ctx *gin.Context) {
	var req listUserRequest
	var userRespList []userResponse
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	// TODO: REFACTOR THIS PART OF CODE SINCE IT IS MESSY
	users, err := server.store.ListUsers(ctx, arg)
	for _, user := range users {
		rsp := newUserResponse(user)
		userRespList = append(userRespList, rsp)
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, userRespList)
}

type updateUserRequestUri struct {
	Uuid int64 `uri:"id" binding:"required,min=1"`
}
type updateUserRequestJson struct {
	FirstName  string  `json:"first_name" binding:"required"`
	MiddleName string  `json:"middle_name" binding:"required"`
	LastName   string  `json:"last_name" binding:"required"`
	Gender     string  `json:"gender" binding:"required,oneof=M F"`
	Age        int16   `json:"age" binding:"required"`
	Balance    float32 `json:"balance" binding:"required"`
	Password   string  `json:"password" binding:"required"`
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
	hashedPassword, err := util.HashPassword(reqJson.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.UpdateUserParams{
		Uuid:           reqUri.Uuid,
		FirstName:      reqJson.FirstName,
		MiddleName:     reqJson.MiddleName,
		LastName:       reqJson.LastName,
		Gender:         reqJson.Gender,
		Age:            reqJson.Age,
		Balance:        reqJson.Balance,
		HashedPassword: hashedPassword,
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
	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
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
	// authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	// if user.Username != authPayload.Username {
	// 	err := errors.New("can't buy on a behalf of other user")
	// 	ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	// 	return
	// }
	r, err := server.store.DeleteUser(ctx, req.Uuid)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if r == 0 {
		ctx.JSON(http.StatusNotFound, notFoundResponse("user"))
		return
	}
	ctx.JSON(http.StatusOK, successDeleteResponse())
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginUserResponse struct {
	AccessToken string       `json:"access_token" binding:"required"`
	User        userResponse `json:"user" binding:"required"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user, err := server.store.GetUserByUserName(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	// Only when the password is correct we will create a new access token
	accessToken, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)

}
