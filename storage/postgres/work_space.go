package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"gitlab.com/samandar_tukhtayev/trello/models"
)

func (r *postgresRepo) WorkSpaceCreate(ctx context.Context, req *models.WorkspaceCreateReq) (*models.WorkspaceResponse, error) {

	res := &models.WorkspaceResponse{}
	query := r.Db.Builder.Insert("work_spaces").Columns(
		"title",
		"owner_id",
	).Values(req.Title, req.OwnerID).Suffix(
		"RETURNING id, title, owner_id, created_at, updated_at")

	err := query.RunWith(r.Db.Db).Scan(
		&res.ID, &res.Title,
		&res.OwnerID,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *TemplateRepo) Create()")
	}
	res.CreatedAt = CreatedAt
	res.UpdatedAt = UpdatedAt

	return res, nil
}

func (r *postgresRepo) WorkSpaceFind(ctx context.Context, req *models.WorkspaceFindReq) (*models.WorkspaceFindResponse, error) {
	var (
		res = &models.WorkspaceFindResponse{}
	)

	countQuery := r.Db.Builder.Select("count(1) as count").From("work_spaces").Where("deleted_at is null")
	err := countQuery.RunWith(r.Db.Db).QueryRow().Scan(&res.Count)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) FindList()")

	}

	query := r.Db.Builder.Select("id, title, owner_id, created_at, updated_at").
		From("work_spaces").Where("deleted_at is null").OrderBy("created_at").Limit(uint64(req.Limit)).Offset(uint64((req.Page - 1) * req.Limit))

	rows, err := query.RunWith(r.Db.Db).Query()
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) FindList()")
	}
	defer rows.Close()

	for rows.Next() {
		workSpace := &models.WorkspaceResponse{}
		err := rows.Scan(
			&workSpace.ID, &workSpace.Title,
			&workSpace.OwnerID,
			&CreatedAt, &UpdatedAt,
		)
		if err != nil {
			return res, HandleDatabaseError(err, r.Log, "(r *models.WorkSpacesWorkSpaceRepo) FindList()")
		}

		workSpace.CreatedAt = CreatedAt
		workSpace.UpdatedAt = UpdatedAt
		res.Workspaces = append(res.Workspaces, workSpace)
	}

	return res, nil
}

func (r *postgresRepo) WorkSpaceGet(ctx context.Context, req *models.WorkspaceGetReq) (*models.WorkspaceResponse, error) {
	query := r.Db.Builder.Select("id, title, owner_id, created_at, updated_at").
		From("work_spaces").
		Where(squirrel.Eq{"id": req.Id})

	res := &models.WorkspaceResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.ID, &res.Title,
		&res.OwnerID,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *TemplateRepo) Get()")
	}

	res.CreatedAt = CreatedAt
	res.UpdatedAt = UpdatedAt

	return res, nil
}

func (r *postgresRepo) WorkSpaceUpdate(ctx context.Context, req *models.WorkspaceUpdateReq) (*models.WorkspaceResponse, error) {
	mp := make(map[string]interface{})
	mp["title"] = req.Title
	mp["owner_id"] = req.OwnerID
	mp["updated_at"] = time.Now()
	query := r.Db.Builder.Update("work_spaces").SetMap(mp).
		Where(squirrel.Eq{"id": req.Id}).
		Suffix("RETURNING id, title, owner_id, created_at, updated_at")

	res := &models.WorkspaceResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.ID, &res.Title,
		&res.OwnerID,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *models.WorkSpaceTemplateRepo) Update()")
	}
	res.CreatedAt = CreatedAt
	res.UpdatedAt = UpdatedAt

	return res, nil
}

func (r *postgresRepo) WorkSpaceDelete(ctx context.Context, req *models.WorkspaceDeleteReq) error {
	query := r.Db.Builder.Delete("work_spaces").Where(squirrel.Eq{"id": req.Id})

	_, err := query.RunWith(r.Db.Db).Exec()
	return HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) Delete()")
}
