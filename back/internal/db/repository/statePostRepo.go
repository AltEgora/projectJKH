package repository

import (
	"ConsultantBack/internal/db/domain"
	"context"
	"database/sql"
	"strings"
)

type StatePostRepo struct {
	db *sql.DB
}

func NewStatePostRepo(db *sql.DB) *StatePostRepo {
	return &StatePostRepo{db: db}
}

func (s *StatePostRepo) GetList(ctx context.Context, category string) ([]domain.State, error) {
	states := []domain.State{}

	query := `
		SELECT * FROM states
		WHERE category=$1
		ORDER BY updated_at
	`

	rows, err := s.db.Query(query, category)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var row domain.State
		rows.Scan(&row.ID, &row.Header, &row.Content, &row.Category, &row.DatePublished, &row.DateUpdated)
		row.DatePublished = strings.Clone(row.DatePublished[:10])
		row.DateUpdated = strings.Clone(row.DateUpdated[:10])

		states = append(states, row)
	}

	return states, nil
}
