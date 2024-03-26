package grpchadlers

import (
	contactsv1 "ContactsService/pkg/pb/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateCreateRequest(req *contactsv1.CreateRequest) error {
	if req.GetFirstName() == "" {
		return status.Errorf(codes.InvalidArgument, "First name is required")
	}
	if req.GetSecondName() == "" {
		return status.Errorf(codes.InvalidArgument, "Second name is required")
	}
	if req.GetMiddleName() == "" {
		return status.Errorf(codes.InvalidArgument, "Middle name is required")
	}
	if req.GetPhoneNumber() == "" {
		return status.Errorf(codes.InvalidArgument, "Phone number is required")
	}
	return nil
}

func ValidateUpdateRequest(req *contactsv1.UpdateRequest) error {
	if req.GetContact().GetId() == "" {
		return status.Errorf(codes.InvalidArgument, "Contact Id is required")
	}
	if req.GetContact().GetFirstName() == "" {
		return status.Errorf(codes.InvalidArgument, "First name is required")
	}
	if req.GetContact().GetSecondName() == "" {
		return status.Errorf(codes.InvalidArgument, "Second name is required")
	}
	if req.GetContact().GetMiddleName() == "" {
		return status.Errorf(codes.InvalidArgument, "Middle name is required")
	}
	if req.GetContact().GetPhoneNumber() == "" {
		return status.Errorf(codes.InvalidArgument, "Phone number is required")
	}
	return nil
}

func ValidateDeleteRequest(req *contactsv1.DeleteRequest) error {
	if req.GetId() == "" {
		return status.Errorf(codes.InvalidArgument, "Id is required")
	}
	return nil
}

func ValidateGetByIDRequest(req *contactsv1.GetByIDRequest) error {
	if req.GetId() == "" {
		return status.Errorf(codes.InvalidArgument, "Id is required")
	}
	return nil
}
