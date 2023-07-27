package postgres

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/ent"
	"github.com/marianozunino/cc-backend-go/ent/contactusmessage"
	"github.com/marianozunino/cc-backend-go/store"
)

var _ store.MessageStore = &MessageStore{}

type MessageStore struct {
	*ent.Client
}

// CountMessages implements store.MessageStore.
func (s *MessageStore) CountMessages(ctx context.Context) (int, error) {
	return s.Client.ContactUsMessage.Query().Where(contactusmessage.DeletedAtIsNil()).Count(ctx)
}

// DeleteMessage implements store.MessageStore.
func (s *MessageStore) DeleteMessage(ctx context.Context, id uuid.UUID) (*ent.ContactUsMessage, error) {
	return s.Client.ContactUsMessage.UpdateOneID(id).Where(contactusmessage.DeletedAtIsNil()).SetDeletedAt(time.Now()).Save(ctx)

}

// Message implements store.MessageStore.
func (s *MessageStore) Message(ctx context.Context, id uuid.UUID) (*ent.ContactUsMessage, error) {
	return s.Client.ContactUsMessage.Query().Where(contactusmessage.ID(id)).Where(contactusmessage.DeletedAtIsNil()).First(ctx)

}

// PaginatedMessages implements store.MessageStore.
func (s *MessageStore) PaginatedMessages(ctx context.Context, orderBy *dtos.MessageAggregationInput, take *int, skip *int) (ent.ContactUsMessages, error) {
	query := s.Client.ContactUsMessage.Query().Where(contactusmessage.DeletedAtIsNil())

	order := []contactusmessage.OrderOption{}

	if orderBy != nil {
		if orderBy.ID != nil {
			if *orderBy.ID == dtos.SortOrderAsc {
				order = append(order, ent.Asc(contactusmessage.FieldID))
			} else {
				order = append(order, ent.Desc(contactusmessage.FieldID))
			}
		}
		if orderBy.Name != nil {
			if *orderBy.Name == dtos.SortOrderAsc {
				order = append(order, ent.Asc(contactusmessage.FieldName))
			} else {
				order = append(order, ent.Desc(contactusmessage.FieldName))
			}
		}
		if orderBy.Email != nil {
			if *orderBy.Email == dtos.SortOrderAsc {
				order = append(order, ent.Asc(contactusmessage.FieldEmail))
			} else {
				order = append(order, ent.Desc(contactusmessage.FieldEmail))
			}
		}
		if orderBy.Phone != nil {
			if *orderBy.Phone == dtos.SortOrderAsc {
				order = append(order, ent.Asc(contactusmessage.FieldPhone))
			} else {
				order = append(order, ent.Desc(contactusmessage.FieldPhone))
			}
		}
	}

	if take != nil {
		query = query.Limit(*take)
	}
	if skip != nil {
		query = query.Offset(*skip)
	}

	if len(order) > 0 {
		query = query.Order(order...)
	}

	return query.All(ctx)
}

// SendMessage implements store.MessageStore.
func (s *MessageStore) SendMessage(ctx context.Context, message *ent.ContactUsMessage) (*ent.ContactUsMessage, error) {
	return s.Client.ContactUsMessage.Create().SetID(uuid.New()).SetName(message.Name).SetEmail(message.Email).SetNillablePhone(message.Phone).SetMessage(message.Message).Save(ctx)
}

// UpdateMessage implements store.MessageStore.
func (s *MessageStore) UpdateMessage(ctx context.Context, message *ent.ContactUsMessage) (*ent.ContactUsMessage, error) {
	return s.Client.ContactUsMessage.UpdateOneID(message.ID).SetName(message.Name).SetEmail(message.Email).SetOrClearPhone(message.Phone).SetMessage(message.Message).Save(ctx)
}
