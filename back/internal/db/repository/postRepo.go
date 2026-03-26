package repository

import (
	"ConsultantBack/internal/db/domain"
	"context"
	"database/sql"
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
		ORDER BY DateUpdated
		OFFSET $1 LIMIT $2
	`

	rows, err := r.db.Query(query, start, count)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var row domain.New
		rows.Scan(row.ID, row.Header, row.URL, row.DatePublished, row.DateUpdated)

		news = append(news, row)
	}

	return news, nil
}

func (r *PostRepo) Create(ctx context.Context, new domain.New) error {
	_, err := r.db.Exec("INSERT INTO News VALUES (default, $1, $2, $3)", new.Header, new.URL, new.DatePublished, new.DateUpdated)
	return err
}

// TODO
func (r *PostRepo) DeleteOld(ctw context.Context, date string) error {
	return nil
}
