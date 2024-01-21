package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/samandar_tukhtayev/trello/models"
)

// @Router		/card [POST]
// @Summary		Create card
// @Tags        Card
// @Description	Here card can be created.
// @Accept      json
// @Produce		json
// @Param       post   body       models.CardCreateReq true "post info"
// @Success		200 	{object}  models.CardApiResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) CardCreate(c *gin.Context) {
	body := &models.CardCreateReq{}
	err := c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.storage.Postgres().CardCreate(context.Background(), body)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().CardCreate()") {
		return
	}

	c.JSON(http.StatusOK, &models.CardApiResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Router		/card/{id} [GET]
// @Summary		Get card by key
// @Tags        Card
// @Description	Here card can be got.
// @Accept      json
// @Produce		json
// @Param       id       path     int true "id"
// @Success		200 	{object}  models.CardApiResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) CardGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi()") {
		return
	}

	res, err := h.storage.Postgres().CardGet(context.Background(), &models.CardGetReq{
		Id: id,
	})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().CardGet()") {
		return
	}

	c.JSON(http.StatusOK, models.CardApiResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Router		/card/list [GET]
// @Summary		Get cards list
// @Tags        Card
// @Description	Here all cards can be got.
// @Accept      json
// @Produce		json
// @Param       filters query models.CardFindReq true "filters"
// @Success		200 	{object}  models.CardApiFindResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) CardFind(c *gin.Context) {
	page, err := ParsePageQueryParam(c)
	if HandleBadRequestErrWithMessage(c, h.log, err, "helper.ParsePageQueryParam(c)") {
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if HandleBadRequestErrWithMessage(c, h.log, err, "helper.ParseLimitQueryParam(c)") {
		return
	}

	res, err := h.storage.Postgres().CardFind(context.Background(), &models.CardFindReq{
		Page:  page,
		Limit: limit,
	})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().CardFind()") {
		return
	}

	c.JSON(http.StatusOK, &models.CardApiFindResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Summary		Update card
// @Tags        Card
// @Description	Here card can be updated.
// @Accept      json
// @Produce		json
// @Param       post   body       models.CardUpdateReq true "post info"
// @Success		200 	{object}  models.CardApiResponse
// @Failure     default {object}  models.DefaultResponse
// @Router		/card [PUT]
func (h *handlerV1) CardUpdate(c *gin.Context) {
	body := &models.CardUpdateReq{}
	err := c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.storage.Postgres().CardUpdate(context.Background(), body)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().CardUpdate()") {
		return
	}

	c.JSON(http.StatusOK, &models.CardApiResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Router		/card/{id} [DELETE]
// @Summary		Delete card
// @Tags        Card
// @Description	Here card can be deleted.
// @Accept      json
// @Produce		json
// @Param       id       path     int true "id"
// @Success		200 	{object}  models.DefaultResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) CardDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi()") {
		return
	}

	err = h.storage.Postgres().CardDelete(context.Background(), &models.CardDeleteReq{Id: id})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().CardDelete()") {
		return
	}

	c.JSON(http.StatusOK, models.DefaultResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "Successfully deleted",
	})
}
