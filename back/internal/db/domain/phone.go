package domain

import (
	"context"
)

type Phone struct {
	ID          int
	Name        string
	Explanation string
	Value       string
}

var PhonePRepo PhoneRepository

type PhoneRepository interface {
	GetList(ctx context.Context) ([]Phone, error)
	Create(ctx context.Context, new Phone) error
	Delete(ctw context.Context, id int) error
}
