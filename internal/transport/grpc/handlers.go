package grpchadlers

import (
	"ContactsService/internal/models/mapper"
	"ContactsService/internal/service"
	contactsv1 "ContactsService/pkg/pb/gen"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type serverAPI struct {
	contactsv1.UnimplementedContactsServiceServer
	cs service.IContactService
}

func RegisterGRPCServer(server *grpc.Server, cs service.IContactService) {
	contactsv1.RegisterContactsServiceServer(server, &serverAPI{cs: cs})
}

func (s *serverAPI) Create(ctx context.Context, req *contactsv1.CreateRequest) (*contactsv1.CreateResponse, error) {
	if err := ValidateCreateRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	id, err := s.cs.CreateContact(ctx, mapper.ContactCreateGRPCtoDB(req))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &contactsv1.CreateResponse{Id: id.String()}, nil
}

func (s *serverAPI) Update(ctx context.Context, req *contactsv1.UpdateRequest) (*contactsv1.UpdateResponse, error) {
	if err := ValidateUpdateRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err := s.cs.UpdateContact(ctx, mapper.ContactUpdateGRPCtoDB(req.GetContact()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &contactsv1.UpdateResponse{Id: req.GetContact().GetId()}, nil
}

func (s *serverAPI) Delete(ctx context.Context, req *contactsv1.DeleteRequest) (*contactsv1.DeleteResponse, error) {
	if err := ValidateDeleteRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	err = s.cs.DeleteContact(ctx, id)
	return &contactsv1.DeleteResponse{Id: req.GetId()}, nil
}

func (s *serverAPI) GetByID(ctx context.Context, req *contactsv1.GetByIDRequest) (*contactsv1.GetByIDResponse, error) {
	if err := ValidateGetByIDRequest(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	id, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	c, err := s.cs.GetContactByID(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &contactsv1.GetByIDResponse{
		Contact: mapper.ContactEntityToGRPC(c),
	}, nil
}

func (s *serverAPI) List(ctx context.Context, req *contactsv1.ListRequest) (*contactsv1.ListResponse, error) {
	contacts, err := s.cs.ListContacts(ctx, nil)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	contactList := make([]*contactsv1.Contact, len(contacts))
	for i, contact := range contacts {
		contactList[i] = mapper.ContactEntityToGRPC(contact)
	}
	return &contactsv1.ListResponse{Contacts: contactList}, nil
}
