package rest

import "github.com/google/uuid"

type CreateContactRequest struct {
	FirstName   string `json:"firstName"`
	SecondName  string `json:"secondName"`
	MiddleName  string `json:"middleName"`
	PhoneNumber string `json:"phoneNumber"`
}

type UpdateContactRequest struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstName"`
	SecondName  string    `json:"secondName"`
	MiddleName  string    `json:"middleName"`
	PhoneNumber string    `json:"phoneNumber"`
}
