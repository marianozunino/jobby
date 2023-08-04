package mapper

import (
	"github.com/marianozunino/jobby/dtos"
	"github.com/marianozunino/jobby/ent"
)

func PostToDto(entity *ent.Post) *dtos.Post {
	dto := &dtos.Post{
		ID:            entity.ID,
		Title:         entity.Title,
		Slug:          entity.Slug,
		PreviewImage:  entity.PreviewImage,
		IsPublished:   entity.IsPublished,
		IsHighlighted: entity.IsHighlighted,
		CreatedAt:     entity.CreatedAt,
		UpdatedAt:     entity.UpdatedAt,
		DeletedAt:     entity.DeletedAt,
	}
	return dto
}

func PostsToDtos(entities ent.Posts) []*dtos.Post {
	dtos := make([]*dtos.Post, len(entities))
	for i, entity := range entities {
		dtos[i] = PostToDto(entity)
	}
	return dtos
}
