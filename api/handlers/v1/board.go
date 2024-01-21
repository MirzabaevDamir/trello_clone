package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/samandar_tukhtayev/trello/models"
)

// @Router		/board [POST]
// @Summary		Create board
// @Tags        Board
// @Description	Here board can be created.
// @Accept      json
// @Produce		json
// @Param       post   body       models.BoardCreateReq true "post info"
// @Success		200 	{object}  models.BoardApiResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) BoardCreate(c *gin.Context) {
	body := &models.BoardCreateReq{}
	err := c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.storage.Postgres().BoardCreate(context.Background(), body)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().BoardCreate()") {
		return
	}

	c.JSON(http.StatusOK, &models.BoardApiResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Router		/board/{id} [GET]
// @Summary		Get board by key
// @Tags        Board
// @Description	Here board can be got.
// @Accept      json
// @Produce		json
// @Param       id       path     int true "id"
// @Success		200 	{object}  models.BoardApiResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) BoardGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi()") {
		return
	}

	res, err := h.storage.Postgres().BoardGet(context.Background(), &models.BoardGetReq{
		Id: id,
	})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().BoardGet()") {
		return
	}

	c.JSON(http.StatusOK, models.BoardApiResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Router		/board/list [GET]
// @Summary		Get boards list
// @Tags        Board
// @Description	Here all boards can be got.
// @Accept      json
// @Produce		json
// @Param       filters query models.BoardFindReq true "filters"
// @Success		200 	{object}  models.BoardApiFindResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) BoardFind(c *gin.Context) {
	page, err := ParsePageQueryParam(c)
	if HandleBadRequestErrWithMessage(c, h.log, err, "helper.ParsePageQueryParam(c)") {
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if HandleBadRequestErrWithMessage(c, h.log, err, "helper.ParseLimitQueryParam(c)") {
		return
	}

	res, err := h.storage.Postgres().BoardFind(context.Background(), &models.BoardFindReq{
		Page:  page,
		Limit: limit,
	})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().BoardFind()") {
		return
	}

	c.JSON(http.StatusOK, &models.BoardApiFindResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Summary		Update board
// @Tags        Board
// @Description	Here board can be updated.
// @Accept      json
// @Produce		json
// @Param       post   body       models.BoardUpdateReq true "post info"
// @Success		200 	{object}  models.BoardApiResponse
// @Failure     default {object}  models.DefaultResponse
// @Router		/board [PUT]
func (h *handlerV1) BoardUpdate(c *gin.Context) {
	body := &models.BoardUpdateReq{}
	err := c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.storage.Postgres().BoardUpdate(context.Background(), body)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().BoardUpdate()") {
		return
	}

	c.JSON(http.StatusOK, &models.BoardApiResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Router		/board/{id} [DELETE]
// @Summary		Delete board
// @Tags        Board
// @Description	Here board can be deleted.
// @Accept      json
// @Produce		json
// @Param       id       path     int true "id"
// @Success		200 	{object}  models.DefaultResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) BoardDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi()") {
		return
	}

	err = h.storage.Postgres().BoardDelete(context.Background(), &models.BoardDeleteReq{Id: id})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().BoardDelete()") {
		return
	}

	c.JSON(http.StatusOK, models.DefaultResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "Successfully deleted",
	})
}
