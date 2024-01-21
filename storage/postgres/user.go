package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"gitlab.com/samandar_tukhtayev/trello/models"
)

func (r *postgresRepo) UserCreate(ctx context.Context, req *models.UserCreateReq) (*models.UserResponse, error) {

	res := &models.UserResponse{}
	query := r.Db.Builder.Insert("users").Columns(
		"first_name",
		"last_name",
		"username",
		"password",
	).Values(req.FirstName, req.LastName, req.Username, req.Password).Suffix(
		"RETURNING id,first_name, last_name, username, created_at, updated_at")

	err := query.RunWith(r.Db.Db).Scan(
		&res.ID, &res.FirstName,
		&res.LastName, &res.Username,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *TemplateRepo) Create()")
	}
	res.CreatedAt = CreatedAt
	res.UpdatedAt = UpdatedAt

	return res, nil
}

func (r *postgresRepo) UserFind(ctx context.Context, req *models.UserFindReq) (*models.UserFindResponse, error) {
	var (
		res = &models.UserFindResponse{}
	)

	countQuery := r.Db.Builder.Select("count(1) as count").From("users").Where("deleted_at is null")
	if req.Search != "" {
		countQuery = countQuery.Where("first_name ILIKE ? OR last_name ILIKE ? OR username ILIKE ?", "%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%")
	}

	err := countQuery.RunWith(r.Db.Db).QueryRow().Scan(&res.Count)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) FindList()")
	}

	query := r.Db.Builder.Select("id, first_name, last_name, username, created_at, updated_at").
		From("users").Where("deleted_at is null")
	if req.Search != "" {
		query = query.Where("first_name ILIKE ? OR last_name ILIKE ? OR username ILIKE ?", "%"+req.Search+"%", "%"+req.Search+"%", "%"+req.Search+"%")
	}
	query = query.OrderBy("created_at").Limit(uint64(req.Limit)).Offset(uint64((req.Page - 1) * req.Limit))

	rows, err := query.RunWith(r.Db.Db).Query()
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) FindList()")
	}
	defer rows.Close()

	for rows.Next() {
		user := &models.UserResponse{}
		err := rows.Scan(
			&user.ID, &user.FirstName,
			&user.LastName, &user.Username,
			&CreatedAt, &UpdatedAt,
		)
		if err != nil {
			return res, HandleDatabaseError(err, r.Log, "(r *models.UsersUserRepo) FindList()")
		}

		user.CreatedAt = CreatedAt
		user.UpdatedAt = UpdatedAt
		res.Users = append(res.Users, user)
	}

	return res, nil
}

func (r *postgresRepo) UserGet(ctx context.Context, req *models.UserGetReq) (*models.UserResponse, error) {
	query := r.Db.Builder.Select("id, first_name, last_name, username, created_at, updated_at").
		From("users").
		Where(squirrel.Eq{"id": req.Id})

	res := &models.UserResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.ID, &res.FirstName,
		&res.LastName, &res.Username,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *TemplateRepo) Get()")
	}

	res.CreatedAt = CreatedAt
	res.UpdatedAt = UpdatedAt

	return res, nil
}

func (r *postgresRepo) UserUpdate(ctx context.Context, req *models.UserUpdateReq) (*models.UserResponse, error) {
	mp := make(map[string]interface{})
	mp["first_name"] = req.FirstName
	mp["last_name"] = req.LastName
	mp["username"] = req.Username
	mp["password"] = req.Password
	mp["updated_at"] = time.Now()
	query := r.Db.Builder.Update("users").SetMap(mp).
		Where(squirrel.Eq{"id": req.Id}).
		Suffix("RETURNING id, first_name, last_name, username, created_at, updated_at")

	res := &models.UserResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.ID, &res.FirstName,
		&res.LastName, &res.Username,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *models.UserTemplateRepo) Update()")
	}
	res.CreatedAt = CreatedAt
	res.UpdatedAt = UpdatedAt

	return res, nil
}

func (r *postgresRepo) UserDelete(ctx context.Context, req *models.UserDeleteReq) error {
	query := r.Db.Builder.Delete("users").Where(squirrel.Eq{"id": req.Id})

	_, err := query.RunWith(r.Db.Db).Exec()
	return HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) Delete()")
}
