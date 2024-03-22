package mapper

import (
	"ContactsService/internal/models/dbmodels"
	"ContactsService/internal/models/entity"
	"ContactsService/internal/models/rest"
	"github.com/google/uuid"
)

func ContactEntityToDB(contact entity.Contact) dbmodels.Contact {
	return dbmodels.Contact{
		ID:          contact.ID,
		FirstName:   contact.FirstName,
		SecondName:  contact.SecondName,
		MiddleName:  contact.MiddleName,
		PhoneNumber: contact.PhoneNumber,
	}
}

func ContactDBToEntity(contact dbmodels.Contact) entity.Contact {
	return entity.Contact{
		ID:          contact.ID,
		FirstName:   contact.FirstName,
		SecondName:  contact.SecondName,
		MiddleName:  contact.MiddleName,
		PhoneNumber: contact.PhoneNumber,
	}
}

func ContactCreateRestToDB(contact rest.CreateContactRequest) dbmodels.Contact {
	return dbmodels.Contact{
		ID:          uuid.New(),
		FirstName:   contact.FirstName,
		SecondName:  contact.SecondName,
		MiddleName:  contact.MiddleName,
		PhoneNumber: contact.PhoneNumber,
	}
}

func ContactUpdateRestToDB(contact rest.UpdateContactRequest) dbmodels.Contact {
	return dbmodels.Contact{
		ID:          contact.ID,
		FirstName:   contact.FirstName,
		SecondName:  contact.SecondName,
		MiddleName:  contact.MiddleName,
		PhoneNumber: contact.PhoneNumber,
	}
}
