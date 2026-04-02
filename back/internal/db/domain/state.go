package domain

import (
	"context"
	"html/template"
)

type State struct {
	ID            int
	Header        string
	Content       template.HTML
	Category      string
	DateUpdated   string
	DatePublished string
}

var StatePRepo StateRepository

type StateRepository interface {
	GetList(ctx context.Context, category string) ([]State, error)
}
