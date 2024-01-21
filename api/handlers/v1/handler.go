package v1

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/samandar_tukhtayev/trello/config"
	"gitlab.com/samandar_tukhtayev/trello/pkg/logger"
	"gitlab.com/samandar_tukhtayev/trello/storage"
)

type HandlerV1I interface {
	UserCreate(c *gin.Context)
	UserGet(c *gin.Context)
	UserFind(c *gin.Context)
	UserUpdate(c *gin.Context)
	UserDelete(c *gin.Context)

	BoardCreate(c *gin.Context)
	BoardGet(c *gin.Context)
	BoardFind(c *gin.Context)
	BoardUpdate(c *gin.Context)
	BoardDelete(c *gin.Context)

	CardCreate(c *gin.Context)
	CardGet(c *gin.Context)
	CardFind(c *gin.Context)
	CardUpdate(c *gin.Context)
	CardDelete(c *gin.Context)

	WorkSpaceCreate(c *gin.Context)
	WorkSpaceGet(c *gin.Context)
	WorkSpaceFind(c *gin.Context)
	WorkSpaceUpdate(c *gin.Context)
	WorkSpaceDelete(c *gin.Context)
}

type handlerV1 struct {
	log     *logger.Logger
	cfg     config.Config
	storage storage.StorageI
}

type HandlerV1Config struct {
	Logger   *logger.Logger
	Cfg      config.Config
	Postgres storage.StorageI
}

// New ...
func New(c *HandlerV1Config) HandlerV1I {
	return &handlerV1{
		log:     c.Logger,
		cfg:     c.Cfg,
		storage: c.Postgres,
	}
}
