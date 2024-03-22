package dbmodels

import "github.com/google/uuid"

type Contact struct {
	ID          uuid.UUID `json:"id"`
	FirstName   string    `json:"firstName"`
	SecondName  string    `json:"secondName"`
	MiddleName  string    `json:"middleName"`
	PhoneNumber string    `json:"phoneNumber"`
}
