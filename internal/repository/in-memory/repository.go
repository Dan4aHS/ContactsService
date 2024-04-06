package in_memory

import (
	"ContactsService/internal/models/dbmodels"
	"ContactsService/internal/models/entity"
	"ContactsService/internal/models/mapper"
	"context"
	"errors"
	"github.com/google/uuid"
	"sync"
)

type ContactRepository struct {
	mx *sync.RWMutex
	db map[uuid.UUID]dbmodels.Contact
}

func NewContactRepository() *ContactRepository {
	return &ContactRepository{mx: new(sync.RWMutex), db: make(map[uuid.UUID]dbmodels.Contact)}
}

func (cr *ContactRepository) CreateContact(_ context.Context, contact entity.Contact) (uuid.UUID, error) {
	c := mapper.ContactEntityToDB(contact)
	cr.mx.RLock()
	defer cr.mx.RUnlock()
	cr.db[c.ID] = c
	return c.ID, nil
}

func (cr *ContactRepository) UpdateContact(_ context.Context, contact entity.Contact) error {
	c := mapper.ContactEntityToDB(contact)
	if _, ok := cr.db[c.ID]; ok {
		cr.db[c.ID] = c
	} else {
		return errors.New("contact does not exist")
	}
	return nil
}

func (cr *ContactRepository) DeleteContact(_ context.Context, id uuid.UUID) error {
	if _, ok := cr.db[id]; ok {
		delete(cr.db, id)
	} else {
		return errors.New("contact does not exist")
	}
	return nil
}

func (cr *ContactRepository) GetContactByID(_ context.Context, id uuid.UUID) (entity.Contact, error) {
	if _, ok := cr.db[id]; ok {
		return mapper.ContactDBToEntity(cr.db[id]), nil
	}
	return entity.Contact{}, errors.New("contact not found")
}

func (cr *ContactRepository) ListContacts(_ context.Context, _ map[string]any) ([]entity.Contact, error) {
	contacts := make([]entity.Contact, 0)
	for _, contact := range cr.db {
		contacts = append(contacts, mapper.ContactDBToEntity(contact))
	}
	return contacts, nil
}

func (cr *ContactRepository) RollBack(_ entity.Contact) error {
	return nil
}

func (cr *ContactRepository) Commit(_ entity.Contact) error {
	return nil
}
