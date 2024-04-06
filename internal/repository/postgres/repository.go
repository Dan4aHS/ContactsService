package postgres

import (
	"ContactsService/internal/models/entity"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ContactRepository struct {
	db *sqlx.DB
	ts map[uuid.UUID]*sqlx.Tx
}

func NewContactRepository(db *sqlx.DB) *ContactRepository {
	return &ContactRepository{db: db, ts: make(map[uuid.UUID]*sqlx.Tx)}
}

func (cr *ContactRepository) CreateContact(ctx context.Context, contact entity.Contact) (uuid.UUID, error) {
	q := `
		INSERT INTO contacts 
		    (id, first_name, second_name, middle_name, phone_number)
		VALUES 
		    ($1, $2, $3, $4, $5)
`
	var err error
	cr.ts[contact.ID], err = cr.db.BeginTxx(ctx, nil)
	if err != nil {
		return uuid.Nil, err
	}
	_, err = cr.ts[contact.ID].ExecContext(ctx, q, contact.ID, contact.FirstName, contact.SecondName, contact.MiddleName, contact.PhoneNumber)
	if err != nil {
		return uuid.Nil, err
	}

	return contact.ID, nil
}

func (cr *ContactRepository) UpdateContact(ctx context.Context, contact entity.Contact) error {
	q := `
		UPDATE 
		    contacts
		SET 
		    first_name = $2,
		    second_name = $3,
		    middle_name = $4,
		    phone_number = $5
		WHERE 
		    id = $1
`
	_, err := cr.db.ExecContext(ctx, q, contact.ID, contact.FirstName, contact.SecondName, contact.MiddleName, contact.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (cr *ContactRepository) DeleteContact(ctx context.Context, id uuid.UUID) error {
	q := `
		DELETE 
		FROM 
		    contacts
		WHERE 
		    id = $1
`
	_, err := cr.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}
	return nil
}

func (cr *ContactRepository) GetContactByID(ctx context.Context, id uuid.UUID) (entity.Contact, error) {
	q := `
		SELECT
		    id, first_name, second_name, middle_name, phone_number
		FROM 
		    contacts
		WHERE
		    id = $1
`
	var contact entity.Contact
	err := cr.db.GetContext(ctx, &contact, q, id)
	if err != nil {
		return contact, err
	}
	return contact, nil
}

func (cr *ContactRepository) ListContacts(ctx context.Context, f map[string]any) ([]entity.Contact, error) {
	var contacts []entity.Contact
	params := squirrel.And{}
	for key, value := range f {
		params = append(params, squirrel.Eq{key: value})
	}
	qf := squirrel.Select("id, first_name, second_name, middle_name, phone_number").
		From("contacts").
		Where(params).
		PlaceholderFormat(squirrel.Dollar)
	q, args, err := qf.ToSql()
	if err != nil {
		return contacts, err
	}
	if len(args) == 0 {
		err = cr.db.SelectContext(ctx, &contacts, q)
	} else {
		err = cr.db.SelectContext(ctx, &contacts, q, args...)
	}
	if err != nil {
		return nil, err
	}
	return contacts, nil
}

func (cr *ContactRepository) RollBack(contact entity.Contact) error {
	if t, ok := cr.ts[contact.ID]; ok {
		if err := t.Rollback(); err != nil {
			return err
		} else {
			delete(cr.ts, contact.ID)
		}
	}
	return nil
}

func (cr *ContactRepository) Commit(contact entity.Contact) error {
	if err := cr.ts[contact.ID].Commit(); err != nil {
		return err
	} else {
		delete(cr.ts, contact.ID)
	}
	return nil
}
