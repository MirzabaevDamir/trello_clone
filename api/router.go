package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golanguzb70/middleware/gin/basicauth"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "gitlab.com/samandar_tukhtayev/trello/api/docs" // docs
	v1 "gitlab.com/samandar_tukhtayev/trello/api/handlers/v1"
	"gitlab.com/samandar_tukhtayev/trello/config"
	"gitlab.com/samandar_tukhtayev/trello/pkg/logger"
	"gitlab.com/samandar_tukhtayev/trello/storage"
)

// Option ...
type Option struct {
	Conf     config.Config
	Logger   *logger.Logger
	Postgres storage.StorageI
}

// New ...
// @title           Trello project API Endpoints
// @version         1.0
// @description     Here QA can test and frontend or mobile developers can get information of API endpoints.

// @BasePath  /v1

// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
func New(log *logger.Logger, cfg config.Config, strg storage.StorageI) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	h := v1.New(&v1.HandlerV1Config{
		Logger:   log,
		Cfg:      cfg,
		Postgres: strg,
	})

	corConfig := cors.DefaultConfig()
	corConfig.AllowAllOrigins = true
	corConfig.AllowCredentials = true
	corConfig.AllowHeaders = []string{"*"}
	corConfig.AllowBrowserExtensions = true
	corConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corConfig))

	authConf := basicauth.Config{
		Users: []basicauth.User{
			{
				UserName: cfg.AdminUsername,
				Password: cfg.AdminPassword,
			},
		},
		RequireAuthForAll: true,
	}

	router.Use(basicauth.New(&authConf).Middleware)
	api := router.Group("/v1")

	user := api.Group("/user")
	user.POST("", h.UserCreate)
	user.GET("/:id", h.UserGet)
	user.GET("/list", h.UserFind)
	user.PUT("", h.UserUpdate)
	user.DELETE(":id", h.UserDelete)

	workSpace := api.Group("/work_space")
	workSpace.POST("", h.WorkSpaceCreate)
	workSpace.GET("/:id", h.WorkSpaceGet)
	workSpace.GET("/list", h.WorkSpaceFind)
	workSpace.PUT("", h.WorkSpaceUpdate)
	workSpace.DELETE(":id", h.WorkSpaceDelete)

	board := api.Group("/board")
	board.POST("", h.BoardCreate)
	board.GET("/:id", h.BoardGet)
	board.GET("/list", h.BoardFind)
	board.PUT("", h.BoardUpdate)
	board.DELETE(":id", h.BoardDelete)

	card := api.Group("/card")
	card.POST("", h.CardCreate)
	card.GET("/:id", h.CardGet)
	card.GET("/list", h.CardFind)
	card.PUT("", h.CardUpdate)
	card.DELETE(":id", h.CardDelete)

	// Don't delete this line, it is used to modify the file automatically

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
