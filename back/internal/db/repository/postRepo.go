package repository

import (
	"ConsultantBack/internal/db/domain"
	"context"
	"database/sql"
	"strings"
)

type PostRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) *PostRepo {
	return &PostRepo{db: db}
}

func (r *PostRepo) GetList(ctx context.Context, start int, count int) ([]domain.New, error) {
	news := []domain.New{}

	query := `
		SELECT * FROM News
		ORDER BY updated_at
		OFFSET $1 LIMIT $2
	`

	rows, err := r.db.Query(query, start, count)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var row domain.New
		rows.Scan(&row.ID, &row.Header, &row.Content, &row.Is_emergency, &row.DatePublished, &row.DateUpdated, &row.Preview)
		row.DatePublished = strings.Clone(row.DatePublished[:10])
		row.DateUpdated = strings.Clone(row.DateUpdated[:10])

		news = append(news, row)
	}

	return news, nil
}

func (r *PostRepo) Create(ctx context.Context, new domain.New) error {
	_, err := r.db.Exec("INSERT INTO News VALUES (default, $1, $2, $3, $4, $5, $6)",
		new.Header, new.Content, new.Is_emergency, new.DatePublished, new.DateUpdated, new.Preview)
	return err
}

func (r *PostRepo) GetShortList(ctx context.Context, start int, count int) ([]domain.ShortNew, error) {
	news := []domain.ShortNew{}

	query := `
		SELECT id, header, is_emergency, updated_at, preview FROM News
		ORDER BY updated_at
		OFFSET $1 LIMIT $2
	`

	rows, err := r.db.Query(query, start, count)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var row domain.ShortNew
		rows.Scan(&row.ID, &row.Header, &row.Is_emergency, &row.DateUpdated, &row.Preview)
		row.DateUpdated = strings.Clone(row.DateUpdated[:10])

		news = append(news, row)
	}

	return news, nil
}

// TODO
func (r *PostRepo) DeleteOld(ctw context.Context, date string) error {
	return nil
}
