package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"gitlab.com/samandar_tukhtayev/trello/models"
)

func (r *postgresRepo) BoardCreate(ctx context.Context, req *models.BoardCreateReq) (*models.BoardResponse, error) {

	res := &models.BoardResponse{}
	query := r.Db.Builder.Insert("boards").Columns(
		"title",
		"order_number",
		"work_space_id",
	).Values(req.Title, req.OrderNumber, req.WorkspaceID).Suffix(
		"RETURNING id,title, order_number, work_space_id, created_at, updated_at")

	err := query.RunWith(r.Db.Db).Scan(
		&res.ID, &res.Title,
		&res.OrderNumber, &res.WorkspaceID,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *TemplateRepo) Create()")
	}
	res.CreatedAt = CreatedAt
	res.UpdatedAt = UpdatedAt

	return res, nil
}

func (r *postgresRepo) BoardFind(ctx context.Context, req *models.BoardFindReq) (*models.BoardFindResponse, error) {
	var (
		res = &models.BoardFindResponse{}
	)

	countQuery := r.Db.Builder.Select("count(1) as count").From("boards").Where("deleted_at is null")
	err := countQuery.RunWith(r.Db.Db).QueryRow().Scan(&res.Count)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) FindList()")

	}

	query := r.Db.Builder.Select("id,title, order_number, work_space_id, created_at, updated_at").
		From("boards").Where("deleted_at is null").OrderBy("created_at").Limit(uint64(req.Limit)).Offset(uint64((req.Page - 1) * req.Limit))

	rows, err := query.RunWith(r.Db.Db).Query()
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) FindList()")
	}
	defer rows.Close()

	for rows.Next() {
		board := &models.BoardResponse{}
		err := rows.Scan(
			&board.ID, &board.Title,
			&board.OrderNumber, &board.WorkspaceID,
			&CreatedAt, &UpdatedAt,
		)
		if err != nil {
			return res, HandleDatabaseError(err, r.Log, "(r *models.BoardsBoardRepo) FindList()")
		}

		board.CreatedAt = CreatedAt
		board.UpdatedAt = UpdatedAt
		res.Boards = append(res.Boards, board)
	}

	return res, nil
}

func (r *postgresRepo) BoardGet(ctx context.Context, req *models.BoardGetReq) (*models.BoardResponse, error) {
	query := r.Db.Builder.Select("id,title, order_number, work_space_id, created_at, updated_at").
		From("boards").
		Where(squirrel.Eq{"id": req.Id})

	res := &models.BoardResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.ID, &res.Title,
		&res.OrderNumber, &res.WorkspaceID,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *TemplateRepo) Get()")
	}

	res.CreatedAt = CreatedAt
	res.UpdatedAt = UpdatedAt

	return res, nil
}

func (r *postgresRepo) BoardUpdate(ctx context.Context, req *models.BoardUpdateReq) (*models.BoardResponse, error) {
	mp := make(map[string]interface{})
	mp["title"] = req.Title
	mp["order_number"] = req.OrderNumber
	mp["work_space_id"] = req.WorkspaceID
	mp["updated_at"] = time.Now()
	query := r.Db.Builder.Update("boards").SetMap(mp).
		Where(squirrel.Eq{"id": req.Id}).
		Suffix("RETURNING id,title, order_number, work_space_id, created_at, updated_at")

	res := &models.BoardResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.ID, &res.Title,
		&res.OrderNumber, &res.WorkspaceID,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *models.BoardTemplateRepo) Update()")
	}
	res.CreatedAt = CreatedAt
	res.UpdatedAt = UpdatedAt

	return res, nil
}

func (r *postgresRepo) BoardDelete(ctx context.Context, req *models.BoardDeleteReq) error {
	query := r.Db.Builder.Delete("boards").Where(squirrel.Eq{"id": req.Id})

	_, err := query.RunWith(r.Db.Db).Exec()
	return HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) Delete()")
}
