package repository

import (
	"ContactsService/internal/models/entity"
	"context"
	"github.com/google/uuid"
)

type IContactRepository interface {
	CreateContact(ctx context.Context, contact entity.Contact) (uuid.UUID, error)
	UpdateContact(ctx context.Context, contact entity.Contact) error
	DeleteContact(ctx context.Context, id uuid.UUID) error
	GetContactByID(ctx context.Context, id uuid.UUID) (entity.Contact, error)
	ListContacts(ctx context.Context, f map[string]any) ([]entity.Contact, error)
	RollBack(contact entity.Contact) error
	Commit(contact entity.Contact) error
}
