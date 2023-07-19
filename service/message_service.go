package service

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/store"
	"github.com/marianozunino/cc-backend-go/store/models"
)

type messageService struct {
	Store store.Store
}

// SendMessage implements MessageService.
func (m *messageService) SendMessage(input dtos.MessageCreateInput) (*dtos.Message, error) {
	message := models.ContactUsMessages{
		Name:    input.Name,
		Email:   input.Email,
		Phone:   sql.NullString{String: input.Phone, Valid: input.Phone != ""},
		Message: input.Message,
	}
	output, err := m.Store.SendMessage(message)
	if err != nil {
		return nil, err
	}
	return m.BuildFromEntity(output), nil
}

// UpdateMessage implements MessageService.
func (m *messageService) UpdateMessage(id string, input dtos.MessageUpdateInput) (*dtos.Message, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	message := models.ContactUsMessages{
		ID: parsedId,
	}
	if input.Name != nil {
		message.Name = *input.Name
	}

	if input.Email != nil {
		message.Email = *input.Email
	}

	if input.Phone != nil {
		message.Phone = sql.NullString{String: *input.Phone, Valid: *input.Phone != ""}
	}

	output, err := m.Store.UpdateMessage(message)
	if err != nil {
		return nil, err
	}
	return m.BuildFromEntity(output), nil
}

// DeleteMessage implements MessageService.
func (m *messageService) DeleteMessage(id string) (*dtos.Message, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	output, err := m.Store.DeleteMessage(parsedId)
	if err != nil {
		return nil, err
	}
	return m.BuildFromEntity(output), nil

}

// GetMessage implements MessageService.
func (m *messageService) GetMessage(id string) (*dtos.Message, error) {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	output, err := m.Store.Message(parsedId)
	if err != nil {
		return nil, err
	}
	return m.BuildFromEntity(output), nil
}

// PaginatedMessages implements MessageService.
func (m *messageService) PaginatedMessages(orderBy *dtos.MessageAggregationInput, take *int, skip *int) (*dtos.PaginatedMessageResponse, error) {
	output, err := m.Store.PaginatedMessages(orderBy, take, skip)
	if err != nil {
		return nil, err
	}
	return &dtos.PaginatedMessageResponse{
		Total: len(output),
		Edges: m.BuildFromEntities(output),
	}, nil

}

func (s *messageService) BuildFromEntity(entity *models.ContactUsMessages) *dtos.Message {
	dto := &dtos.Message{
		ID:        entity.ID.String(),
		Name:      entity.Name,
		Email:     entity.Email,
		Phone:     entity.Phone.String,
		Message:   entity.Message,
		CreatedAt: entity.CreatedAt.Time,
		UpdatedAt: entity.UpdatedAt.Time,
		DeletedAt: getTimeOrNil(entity.DeletedAt),
	}
	return dto
}

func (s *messageService) BuildFromEntities(entities []*models.ContactUsMessages) []dtos.Message {
	dtos := make([]dtos.Message, len(entities))
	for i, entity := range entities {
		dtos[i] = *s.BuildFromEntity(entity)
	}
	return dtos
}

var _ MessageService = &messageService{}
