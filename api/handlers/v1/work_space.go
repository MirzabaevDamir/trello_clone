package v1

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gitlab.com/samandar_tukhtayev/trello/models"
)

// @Router		/work_space [POST]
// @Summary		Create workSpace
// @Tags        WorkSpace
// @Description	Here workSpace can be created.
// @Accept      json
// @Produce		json
// @Param       post   body       models.WorkspaceCreateReq true "post info"
// @Success		200 	{object}  models.WorkspaceApiResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) WorkSpaceCreate(c *gin.Context) {
	body := &models.WorkspaceCreateReq{}
	err := c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.storage.Postgres().WorkSpaceCreate(context.Background(), body)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().WorkSpaceCreate()") {
		return
	}

	c.JSON(http.StatusOK, &models.WorkspaceApiResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Router		/work_space/{id} [GET]
// @Summary		Get workSpace by key
// @Tags        WorkSpace
// @Description	Here workSpace can be got.
// @Accept      json
// @Produce		json
// @Param       id       path     int true "id"
// @Success		200 	{object}  models.WorkspaceApiResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) WorkSpaceGet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi()") {
		return
	}

	res, err := h.storage.Postgres().WorkSpaceGet(context.Background(), &models.WorkspaceGetReq{
		Id: id,
	})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().WorkSpaceGet()") {
		return
	}

	c.JSON(http.StatusOK, models.WorkspaceApiResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Router		/work_space/list [GET]
// @Summary		Get workSpaces list
// @Tags        WorkSpace
// @Description	Here all workSpaces can be got.
// @Accept      json
// @Produce		json
// @Param       filters query models.WorkspaceFindReq true "filters"
// @Success		200 	{object}  models.WorkspaceApiFindResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) WorkSpaceFind(c *gin.Context) {
	page, err := ParsePageQueryParam(c)
	if HandleBadRequestErrWithMessage(c, h.log, err, "helper.ParsePageQueryParam(c)") {
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if HandleBadRequestErrWithMessage(c, h.log, err, "helper.ParseLimitQueryParam(c)") {
		return
	}

	res, err := h.storage.Postgres().WorkSpaceFind(context.Background(), &models.WorkspaceFindReq{
		Page:  page,
		Limit: limit,
	})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().WorkSpaceFind()") {
		return
	}

	c.JSON(http.StatusOK, &models.WorkspaceApiFindResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Summary		Update workSpace
// @Tags        WorkSpace
// @Description	Here workSpace can be updated.
// @Accept      json
// @Produce		json
// @Param       post   body       models.WorkspaceUpdateReq true "post info"
// @Success		200 	{object}  models.WorkspaceApiResponse
// @Failure     default {object}  models.DefaultResponse
// @Router		/work_space [PUT]
func (h *handlerV1) WorkSpaceUpdate(c *gin.Context) {
	body := &models.WorkspaceUpdateReq{}
	err := c.ShouldBindJSON(&body)
	if HandleBadRequestErrWithMessage(c, h.log, err, "c.ShouldBindJSON(&body)") {
		return
	}

	res, err := h.storage.Postgres().WorkSpaceUpdate(context.Background(), body)
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().WorkSpaceUpdate()") {
		return
	}

	c.JSON(http.StatusOK, &models.WorkspaceApiResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "",
		Body:         res,
	})
}

// @Router		/work_space/{id} [DELETE]
// @Summary		Delete workSpace
// @Tags        WorkSpace
// @Description	Here workSpace can be deleted.
// @Accept      json
// @Produce		json
// @Param       id       path     int true "id"
// @Success		200 	{object}  models.DefaultResponse
// @Failure     default {object}  models.DefaultResponse
func (h *handlerV1) WorkSpaceDelete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if HandleBadRequestErrWithMessage(c, h.log, err, "strconv.Atoi()") {
		return
	}

	err = h.storage.Postgres().WorkSpaceDelete(context.Background(), &models.WorkspaceDeleteReq{Id: id})
	if HandleDatabaseLevelWithMessage(c, h.log, err, "h.storage.Postgres().WorkSpaceDelete()") {
		return
	}

	c.JSON(http.StatusOK, models.DefaultResponse{
		ErrorCode:    ErrorSuccessCode,
		ErrorMessage: "Successfully deleted",
	})
}
