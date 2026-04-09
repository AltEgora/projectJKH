package repository

import (
	"ConsultantBack/internal/db/domain"
	"context"
	"database/sql"
)

type PhonePostRepo struct {
	db *sql.DB
}

func NewPhonePostRepo(db *sql.DB) *PhonePostRepo {
	return &PhonePostRepo{db: db}
}

func (r *PhonePostRepo) GetList(ctx context.Context) ([]domain.Phone, error) {
	news := []domain.Phone{}

	query := `
		SELECT * FROM phones
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var row domain.Phone
		rows.Scan(&row.ID, &row.Name, &row.Explanation, &row.Value)

		news = append(news, row)
	}

	return news, nil
}

func (r *PhonePostRepo) Create(ctx context.Context, new domain.Phone) error {
	_, err := r.db.Exec("INSERT INTO News VALUES (default, $1, $2, $3,)",
		new.Name, new.Explanation, new.Value)
	return err
}

// TODO
func (r *PhonePostRepo) Delete(ctw context.Context, id int) error {
	return nil
}
