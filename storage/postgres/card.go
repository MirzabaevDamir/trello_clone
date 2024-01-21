package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.com/samandar_tukhtayev/trello/models"
)

func (r *postgresRepo) CardCreate(ctx context.Context, req *models.CardCreateReq) (*models.CardResponse, error) {
	res := &models.CardResponse{}

	query := r.Db.Builder.
		Insert("cards").
		Columns("title", "description", "user_ids", "deadline", "order_number", "board_id").
		Values(req.Title, req.Description, pq.Array(req.UserIDs), req.Deadline, req.OrderNumber, req.BoardID).
		Suffix("RETURNING id, title, description, user_ids, deadline, order_number, board_id, created_at, updated_at")

	err := query.RunWith(r.Db.Db).Scan(
		&res.ID, &res.Title, &res.Description, pq.Array(&res.UserIDs), &res.Deadline, &res.OrderNumber, &res.BoardID, &res.CreatedAt, &res.UpdatedAt,
	)
	if err != nil {
		return nil, HandleDatabaseError(err, r.Log, "(r *postgresRepo) CardCreate()")
	}

	return res, nil
}

func (r *postgresRepo) CardFind(ctx context.Context, req *models.CardFindReq) (*models.CardFindResponse, error) {
	var (
		res = &models.CardFindResponse{}
	)

	countQuery := r.Db.Builder.Select("count(1) as count").From("cards").Where("deleted_at is null")
	err := countQuery.RunWith(r.Db.Db).QueryRow().Scan(&res.Count)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) FindList()")

	}

	query := r.Db.Builder.Select("id, title, description, user_ids, deadline, order_number, board_id, created_at, updated_at").
		From("cards").Where("deleted_at is null").OrderBy("created_at").Limit(uint64(req.Limit)).Offset(uint64((req.Page - 1) * req.Limit))

	rows, err := query.RunWith(r.Db.Db).Query()
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) FindList()")
	}
	defer rows.Close()

	for rows.Next() {
		card := &models.CardResponse{}
		var userIDs pq.Int64Array
		err := rows.Scan(
			&card.ID, &card.Title,
			&card.Description, &userIDs,
			&card.Deadline, &card.OrderNumber,
			&card.BoardID,
			&CreatedAt, &UpdatedAt,
		)
		if err != nil {
			return res, HandleDatabaseError(err, r.Log, "(r *models.CardsCardRepo) FindList()")
		}
		card.UserIDs = userIDs

		card.CreatedAt = CreatedAt
		card.UpdatedAt = UpdatedAt
		res.Cards = append(res.Cards, card)
	}

	return res, nil
}

func (r *postgresRepo) CardGet(ctx context.Context, req *models.CardGetReq) (*models.CardResponse, error) {
	query := r.Db.Builder.Select("id, title, description, user_ids, deadline, order_number, board_id, created_at, updated_at").
		From("cards").
		Where(squirrel.Eq{"id": req.Id})
	var userIDs pq.Int64Array

	res := &models.CardResponse{}
	err := query.RunWith(r.Db.Db).QueryRow().Scan(
		&res.ID, &res.Title,
		&res.Description, &userIDs,
		&res.Deadline, &res.OrderNumber,
		&res.BoardID,
		&CreatedAt, &UpdatedAt,
	)
	if err != nil {
		return res, HandleDatabaseError(err, r.Log, "(r *TemplateRepo) Get()")
	}
	res.UserIDs = userIDs
	res.CreatedAt = CreatedAt
	res.UpdatedAt = UpdatedAt

	return res, nil
}

func (r *postgresRepo) CardUpdate(ctx context.Context, req *models.CardUpdateReq) (*models.CardResponse, error) {
	query := `
	UPDATE cards
	SET title = $1, description = $2, user_ids = $3, deadline = $4, order_number = $5, board_id = $6, updated_at = $7
	WHERE id = $8
	RETURNING id, title, description, user_ids, deadline, order_number, board_id, created_at, updated_at
	`

	var res models.CardResponse
	var userIDs pq.Int64Array
	var createdAt, updatedAt time.Time

	err := r.Db.Db.QueryRowContext(ctx, query, req.Title, req.Description, pq.Array(req.UserIDs), req.Deadline, req.OrderNumber, req.BoardID, time.Now(), req.ID).
		Scan(&res.ID, &res.Title, &res.Description, &userIDs, &res.Deadline, &res.OrderNumber, &res.BoardID, &createdAt, &updatedAt)

	if err != nil {
		return &res, HandleDatabaseError(err, r.Log, "(r *models.CardTemplateRepo) Update()")
	}

	res.UserIDs = userIDs
	res.CreatedAt = createdAt
	res.UpdatedAt = updatedAt

	return &res, nil
}

func (r *postgresRepo) CardDelete(ctx context.Context, req *models.CardDeleteReq) error {
	query := r.Db.Builder.Delete("cards").Where(squirrel.Eq{"id": req.Id})

	_, err := query.RunWith(r.Db.Db).Exec()
	return HandleDatabaseError(err, r.Log, "(r *models.TemplateTemplateRepo) Delete()")
}
