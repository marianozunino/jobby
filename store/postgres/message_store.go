package postgres

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/marianozunino/cc-backend-go/dtos"
	"github.com/marianozunino/cc-backend-go/store"
	"github.com/marianozunino/cc-backend-go/store/models"
)

var _ store.MessageStore = &MessageStore{}

type MessageStore struct {
	*sqlx.DB
}

// CountMessages implements store.MessageStore.
func (s *MessageStore) CountMessages() (int, error) {
	query := `SELECT COUNT(*) FROM contact_us_messages`
	var count int
	err := s.Get(&count, query)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// DeleteMessage implements store.MessageStore.
func (s *MessageStore) DeleteMessage(id uuid.UUID) (*models.ContactUsMessages, error) {
	query := `DELETE FROM contact_us_messages WHERE id = $1 RETURNING *`
	output := models.ContactUsMessages{}

	err := s.Get(&output, query, id)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// Message implements store.MessageStore.
func (s *MessageStore) Message(id uuid.UUID) (*models.ContactUsMessages, error) {
	query := `SELECT * FROM contact_us_messages WHERE id = $1`
	output := models.ContactUsMessages{}

	err := s.Get(&output, query, id)
	if err != nil {
		return nil, err
	}
	return &output, nil
}

// PaginatedMessages implements store.MessageStore.
func (s *MessageStore) PaginatedMessages(orderBy *dtos.MessageAggregationInput, take *int, skip *int) ([]*models.ContactUsMessages, error) {
	query := "SELECT * FROM contact_us_messages"

	if orderBy != nil {
		query += " ORDER BY "
		var sortFields []string

		switch {
		case orderBy.ID != nil:
			sortFields = append(sortFields, fmt.Sprintf("id %s", orderBy.ID))
		case orderBy.Name != nil:
			sortFields = append(sortFields, fmt.Sprintf("name %s", orderBy.Name))
		case orderBy.Email != nil:
			sortFields = append(sortFields, fmt.Sprintf("email %s", orderBy.Email))
		case orderBy.Phone != nil:
			sortFields = append(sortFields, fmt.Sprintf("phone %s", orderBy.Phone))
		case orderBy.CreatedAt != nil:
			sortFields = append(sortFields, fmt.Sprintf("created_at %s", orderBy.CreatedAt))
		case orderBy.UpdatedAt != nil:
			sortFields = append(sortFields, fmt.Sprintf("updated_at %s", orderBy.UpdatedAt))
		case orderBy.DeletedAt != nil:
			sortFields = append(sortFields, fmt.Sprintf("deleted_at %s", orderBy.DeletedAt))
		}
		query += strings.Join(sortFields, ", ")
	}

	if take != nil {
		query += fmt.Sprintf(" LIMIT %d", *take)
	}
	if skip != nil {
		query += fmt.Sprintf(" OFFSET %d", *skip)
	}

	output := []*models.ContactUsMessages{}
	err := s.Select(&output, query)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// SendMessage implements store.MessageStore.
func (s *MessageStore) SendMessage(message models.ContactUsMessages) (*models.ContactUsMessages, error) {
	query := `INSERT INTO contact_us_messages (name, email, phone, message) VALUES ($1, $2, $3, $4) RETURNING *`
	output := models.ContactUsMessages{}

	err := s.Get(&output, query, message.Name, message.Email, message.Phone, message.Message)
	if err != nil {
		return nil, err
	}
	return &output, nil

}

// UpdateMessage implements store.MessageStore.
func (s *MessageStore) UpdateMessage(message models.ContactUsMessages) (*models.ContactUsMessages, error) {
	query := `UPDATE contact_us_messages SET name = $1, email = $2, phone = $3, message = $4 WHERE id = $5 RETURNING *`
	output := models.ContactUsMessages{}

	err := s.Get(&output, query, message.Name, message.Email, message.Phone, message.Message, message.ID)
	if err != nil {
		return nil, err
	}
	return &output, nil
}
