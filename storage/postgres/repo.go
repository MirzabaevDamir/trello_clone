package postgres

import (
	"context"

	"gitlab.com/samandar_tukhtayev/trello/models"
)

type PostgresI interface {

	UserCreate(ctx context.Context, req *models.UserCreateReq) (*models.UserResponse, error)
	UserGet(ctx context.Context, req *models.UserGetReq) (*models.UserResponse, error)
	UserFind(ctx context.Context, req *models.UserFindReq) (*models.UserFindResponse, error)
	UserUpdate(ctx context.Context, req *models.UserUpdateReq) (*models.UserResponse, error)
	UserDelete(ctx context.Context, req *models.UserDeleteReq) error

	BoardCreate(ctx context.Context, req *models.BoardCreateReq) (*models.BoardResponse, error)
	BoardGet(ctx context.Context, req *models.BoardGetReq) (*models.BoardResponse, error)
	BoardFind(ctx context.Context, req *models.BoardFindReq) (*models.BoardFindResponse, error)
	BoardUpdate(ctx context.Context, req *models.BoardUpdateReq) (*models.BoardResponse, error)
	BoardDelete(ctx context.Context, req *models.BoardDeleteReq) error

	CardCreate(ctx context.Context, req *models.CardCreateReq) (*models.CardResponse, error)
	CardGet(ctx context.Context, req *models.CardGetReq) (*models.CardResponse, error)
	CardFind(ctx context.Context, req *models.CardFindReq) (*models.CardFindResponse, error)
	CardUpdate(ctx context.Context, req *models.CardUpdateReq) (*models.CardResponse, error)
	CardDelete(ctx context.Context, req *models.CardDeleteReq) error

	WorkSpaceCreate(ctx context.Context, req *models.WorkspaceCreateReq) (*models.WorkspaceResponse, error)
	WorkSpaceGet(ctx context.Context, req *models.WorkspaceGetReq) (*models.WorkspaceResponse, error)
	WorkSpaceFind(ctx context.Context, req *models.WorkspaceFindReq) (*models.WorkspaceFindResponse, error)
	WorkSpaceUpdate(ctx context.Context, req *models.WorkspaceUpdateReq) (*models.WorkspaceResponse, error)
	WorkSpaceDelete(ctx context.Context, req *models.WorkspaceDeleteReq) error
}
