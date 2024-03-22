package entity

import "github.com/google/uuid"

type Contact struct {
	ID          uuid.UUID `json:"id" db:"id"`
	FirstName   string    `json:"firstName" db:"first_name"`
	SecondName  string    `json:"secondName" db:"second_name"`
	MiddleName  string    `json:"middleName" db:"middle_name"`
	PhoneNumber string    `json:"phoneNumber" db:"phone_number"`
}
