package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/levietcuong2602/simplebank/db/sqlc"
	"github.com/levietcuong2602/simplebank/utils"
)

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

type GetAccountRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

type GetListAccountsRequest struct {
	PageNum int32 `form:"page_num" binding:"required,min=1"`
	Limit   int32 `form:"limit" binding:"required,min=1"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) GetAccount(ctx *gin.Context) {
	var req GetAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

func (server *Server) GetListAccounts(ctx *gin.Context) {
	var req GetListAccountsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	offset := (req.PageNum - 1) * req.Limit
	arg := db.GetListAccountsParams{
		Limit:  int32(req.Limit),
		Offset: int32(offset),
	}
	accounts, err := server.store.GetListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	data := make([]interface{}, len(accounts))
	for i, v := range accounts {
		data[i] = v
	}

	countAccounts, err := server.store.CountAccounts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	newPaginationParams := utils.NewPaginationParams{
		PageNum:    req.PageNum,
		Limit:      req.Limit,
		TotalCount: int32(countAccounts),
		Data:       data,
	}
	pagination := utils.NewPagination(newPaginationParams)
	ctx.JSON(http.StatusOK, pagination)
}
