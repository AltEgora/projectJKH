package domain

import "context"

type New struct {
	ID            int
	Header        string
	Content       string
	Is_emergency  bool
	DatePublished string
	DateUpdated   string
	Preview       string
}

var PRepo NewRepository

type NewRepository interface {
	GetList(ctx context.Context, start int, count int) ([]New, error)
	Create(ctx context.Context, new New) error
	DeleteOld(ctw context.Context, date string) error
}
