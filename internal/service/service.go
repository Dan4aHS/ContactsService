package service

import (
	"ContactsService/internal/models/dbmodels"
	"ContactsService/internal/models/entity"
	"ContactsService/internal/models/mapper"
	"ContactsService/internal/repository"
	"context"
	"github.com/google/uuid"
)

type ContactService struct {
	Repo repository.IContactRepository `json:"repo"`
}

func NewContactService(repo repository.IContactRepository) *ContactService {
	return &ContactService{Repo: repo}
}

func (cs *ContactService) CreateContact(ctx context.Context, contact dbmodels.Contact) (uuid.UUID, error) {
	return cs.Repo.CreateContact(ctx, mapper.ContactDBToEntity(contact))
}

func (cs *ContactService) UpdateContact(ctx context.Context, contact dbmodels.Contact) error {
	return cs.Repo.UpdateContact(ctx, mapper.ContactDBToEntity(contact))
}

func (cs *ContactService) DeleteContact(ctx context.Context, id uuid.UUID) error {
	return cs.Repo.DeleteContact(ctx, id)
}

func (cs *ContactService) GetContactByID(ctx context.Context, id uuid.UUID) (entity.Contact, error) {
	return cs.Repo.GetContactByID(ctx, id)
}

func (cs *ContactService) ListContacts(ctx context.Context, f map[string]any) ([]entity.Contact, error) {
	return cs.Repo.ListContacts(ctx, f)
}
