package service

import (
	"ContactsService/internal/clients/rabbit-amqp"
	"ContactsService/internal/models/dbmodels"
	"ContactsService/internal/models/entity"
	"ContactsService/internal/models/mapper"
	"ContactsService/internal/repository"
	"context"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"log"
)

type ContactService struct {
	Repo   repository.IContactRepository `json:"repo"`
	Broker *rabbit_amqp.Broker           `json:"broker"`
}

func NewContactService(repo repository.IContactRepository, conn *amqp.Connection) *ContactService {
	broker := rabbit_amqp.NewBroker(conn)
	return &ContactService{Repo: repo, Broker: broker}
}

func (cs *ContactService) CreateContact(ctx context.Context, contact dbmodels.Contact) (uuid.UUID, error) {
	id, err := cs.Repo.CreateContact(ctx, mapper.ContactDBToEntity(contact))
	if err != nil {
		brErr := cs.Broker.SendErrorMessage(err)
		if brErr != nil {
			return uuid.Nil, brErr
		}
		return uuid.Nil, err
	}
	err = cs.Broker.SendContactMessage(contact)
	if err != nil {
		rbErr := cs.Repo.RollBack(mapper.ContactDBToEntity(contact))
		if rbErr != nil {
			return uuid.Nil, rbErr
		}
		return uuid.Nil, err
	} else {
		cErr := cs.Repo.Commit(mapper.ContactDBToEntity(contact))
		if cErr != nil {
			return uuid.Nil, cErr
		}
	}
	return id, nil
}

func (cs *ContactService) UpdateContact(ctx context.Context, contact dbmodels.Contact) error {
	err := cs.Repo.UpdateContact(ctx, mapper.ContactDBToEntity(contact))
	if err != nil {
		brErr := cs.Broker.SendErrorMessage(err)
		if brErr != nil {
			log.Fatal(brErr)
		}
		return err
	}
	return nil
}

func (cs *ContactService) DeleteContact(ctx context.Context, id uuid.UUID) error {
	err := cs.Repo.DeleteContact(ctx, id)
	if err != nil {
		brErr := cs.Broker.SendErrorMessage(err)
		if brErr != nil {
			log.Fatal(brErr)
		}
		return err
	}
	return nil
}

func (cs *ContactService) GetContactByID(ctx context.Context, id uuid.UUID) (entity.Contact, error) {
	c, err := cs.Repo.GetContactByID(ctx, id)
	if err != nil {
		brErr := cs.Broker.SendErrorMessage(err)
		if brErr != nil {
			log.Fatal(brErr)
		}
		return entity.Contact{}, err
	}
	return c, nil
}

func (cs *ContactService) ListContacts(ctx context.Context, f map[string]any) ([]entity.Contact, error) {
	c, err := cs.Repo.ListContacts(ctx, f)
	if err != nil {
		brErr := cs.Broker.SendErrorMessage(err)
		if brErr != nil {
			log.Fatal(brErr)
		}
		return nil, err
	}
	return c, nil
}
