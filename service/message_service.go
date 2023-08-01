package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/ent"
	"github.com/marianozunino/cc-backend-go/store"
)

type messageService struct {
	Store store.Store
}

var _ MessageService = &messageService{}

// SendMessage implements MessageService.
func (m *messageService) SendMessage(ctx context.Context, input dtos.MessageCreateInput) (*dtos.Message, error) {
	message := &ent.ContactUsMessage{
		Name:    input.Name,
		Email:   input.Email,
		Phone:   input.Phone,
		Message: input.Message,
	}
	output, err := m.Store.SendMessage(ctx, message)
	if err != nil {
		return nil, err
	}
	return m.BuildFromEntity(output), nil
}

// UpdateMessage implements MessageService.
func (m *messageService) UpdateMessage(ctx context.Context, id string, input dtos.MessageUpdateInput) (*dtos.Message, error) {
	parsedId, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	message := &ent.ContactUsMessage{}

	message.ID = parsedId

	if input.Name != nil {
		message.Name = *input.Name
	}

	if input.Email != nil {
		message.Email = *input.Email
	}

	message.Phone = input.Phone

	output, err := m.Store.UpdateMessage(ctx, message)
	if err != nil {
		return nil, err
	}
	return m.BuildFromEntity(output), nil
}

// DeleteMessage implements MessageService.
func (m *messageService) DeleteMessage(ctx context.Context, id string) (*dtos.Message, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	output, err := m.Store.DeleteMessage(ctx, parsedId)
	if err != nil {
		return nil, err
	}
	return m.BuildFromEntity(output), nil

}

// GetMessage implements MessageService.
func (m *messageService) GetMessage(ctx context.Context, id string) (*dtos.Message, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	output, err := m.Store.Message(ctx, parsedId)
	if err != nil {
		return nil, err
	}
	return m.BuildFromEntity(output), nil
}

// PaginatedMessages implements MessageService.
func (m *messageService) PaginatedMessages(ctx context.Context, orderBy *dtos.MessageAggregationInput, take *int, skip *int) (*dtos.PaginatedMessageResponse, error) {
	output, err := m.Store.PaginatedMessages(ctx, orderBy, take, skip)
	if err != nil {
		return nil, err
	}
	return &dtos.PaginatedMessageResponse{
		Total: len(output),
		Edges: m.BuildFromEntities(output),
	}, nil

}

func (s *messageService) BuildFromEntity(entity *ent.ContactUsMessage) *dtos.Message {
	dto := &dtos.Message{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Phone:     entity.Phone,
		Message:   entity.Message,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
		DeletedAt: entity.DeletedAt,
	}
	return dto
}

func (s *messageService) BuildFromEntities(entities ent.ContactUsMessages) []*dtos.Message {
	dtos := make([]*dtos.Message, len(entities))
	for i, entity := range entities {
		dtos[i] = s.BuildFromEntity(entity)
	}
	return dtos
}
