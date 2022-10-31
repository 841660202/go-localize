package v1

import (
	"fmt"
	controller "go-localize/src/controller"
	"go-localize/src/model"
	httputil "go-localize/src/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ShowAccount godoc
//
//	@Summary      Show an account
//	@Description  get string by ID
//
// @ID ShowAccount
//
//	@Tags         账号
//	@Accept       json
//	@Produce      json
//	@Param        id   path      int  true  "Account ID"
//	@Success      200  {object}  model.Account
//	@Failure      400  {object}  httputil.HTTPError
//	@Failure      404  {object}  httputil.HTTPError
//	@Failure      500  {object}  httputil.HTTPError
//	@Router       /accounts/{id} [get]
func ShowAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account, err := model.AccountOne(aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, account)
}

// ListAccounts godoc
//
//	@Summary      List accounts
//	@Description  get accounts
//
// @ID ListAccounts
//
//	@Tags         账号
//	@Accept       json
//	@Produce      json
//	@Param        q    query     string  false  "name search by q"  Format(email)
//	@Success      200  {array}   model.Account
//	@Failure      400  {object}  httputil.HTTPError
//	@Failure      404  {object}  httputil.HTTPError
//	@Failure      500  {object}  httputil.HTTPError
//	@Router       /accounts [get]
func ListAccounts(ctx *gin.Context) {
	q := ctx.Request.URL.Query().Get("q")
	accounts, err := model.AccountsAll(q)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

// AddAccount godoc
//
//	@Summary      Add an account
//	@Description  add by json account
//
// @ID AddAccount
//
//	@Tags         账号
//	@Accept       json
//	@Produce      json
//	@Param        account  body      model.AddAccount  true  "Add account"
//	@Success      200      {object}  model.Account
//	@Failure      400      {object}  httputil.HTTPError
//	@Failure      404      {object}  httputil.HTTPError
//	@Failure      500      {object}  httputil.HTTPError
//	@Router       /accounts [post]
func AddAccount(ctx *gin.Context) {
	var addAccount model.AddAccount
	if err := ctx.ShouldBindJSON(&addAccount); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	if err := addAccount.Validation(); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account := model.Account{
		Name: addAccount.Name,
	}
	lastID, err := account.Insert()
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account.ID = lastID
	ctx.JSON(http.StatusOK, account)
}

// UpdateAccount godoc
//
//	@Summary      Update an account
//	@Description  Update by json account
//
// @ID UpdateAccount
//
//	@Tags         账号
//	@Accept       json
//	@Produce      json
//	@Param        id       path      int                  true  "Account ID"
//	@Param        account  body      model.UpdateAccount  true  "Update account"
//	@Success      200      {object}  model.Account
//	@Failure      400      {object}  httputil.HTTPError
//	@Failure      404      {object}  httputil.HTTPError
//	@Failure      500      {object}  httputil.HTTPError
//	@Router       /accounts/{id} [patch]
func UpdateAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var updateAccount model.UpdateAccount
	if err := ctx.ShouldBindJSON(&updateAccount); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	account := model.Account{
		ID:   aid,
		Name: updateAccount.Name,
	}
	err = account.Update()
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, account)
}

// DeleteAccount godoc
//
//	@Summary      Delete an account
//	@Description  Delete by account ID
//
// @ID DeleteAccount
//
//	@Tags         账号
//	@Accept       json
//	@Produce      json
//	@Param        id   path      int  true  "Account ID"  Format(int64)
//	@Success      204  {object}  model.Account
//	@Failure      400  {object}  httputil.HTTPError
//	@Failure      404  {object}  httputil.HTTPError
//	@Failure      500  {object}  httputil.HTTPError
//	@Router       /accounts/{id} [delete]
func DeleteAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err = model.Delete(aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

// UploadAccountImage godoc
//
//	@Summary      Upload account image
//	@Description  Upload file
//
// @ID UploadAccountImage
//
//	@Tags         账号
//	@Accept       multipart/form-data
//	@Produce      json
//	@Param        id    path      int   true  "Account ID"
//	@Param        file  formData  file  true  "account image"
//	@Success      200   {object}  controller.Message
//	@Failure      400   {object}  httputil.HTTPError
//	@Failure      404   {object}  httputil.HTTPError
//	@Failure      500   {object}  httputil.HTTPError
//	@Router       /accounts/{id}/images [post]
func UploadAccountImage(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, controller.Message{Message: fmt.Sprintf("upload complete userID=%d filename=%s", id, file.Filename)})
}
