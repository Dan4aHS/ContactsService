package mapper

import (
	"ContactsService/internal/models/dbmodels"
	"ContactsService/internal/models/entity"
	"ContactsService/internal/models/rest"
	contactsv1 "ContactsService/pkg/pb/gen"
	"fmt"
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

func ContactCreateGRPCtoDB(contact *contactsv1.CreateRequest) dbmodels.Contact {
	return dbmodels.Contact{
		ID:          uuid.New(),
		FirstName:   contact.FirstName,
		SecondName:  contact.SecondName,
		MiddleName:  contact.MiddleName,
		PhoneNumber: contact.PhoneNumber,
	}
}

func ContactUpdateGRPCtoDB(contact *contactsv1.Contact) dbmodels.Contact {
	id, err := uuid.Parse(contact.Id)
	if err != nil {
		fmt.Println(err)
	}
	return dbmodels.Contact{
		ID:          id,
		FirstName:   contact.FirstName,
		SecondName:  contact.SecondName,
		MiddleName:  contact.MiddleName,
		PhoneNumber: contact.PhoneNumber,
	}
}

func ContactEntityToGRPC(contact entity.Contact) *contactsv1.Contact {
	return &contactsv1.Contact{
		Id:          contact.ID.String(),
		FirstName:   contact.FirstName,
		SecondName:  contact.SecondName,
		MiddleName:  contact.MiddleName,
		PhoneNumber: contact.PhoneNumber,
	}
}
