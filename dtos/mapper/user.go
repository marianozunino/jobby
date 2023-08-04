package mapper

import (
	"github.com/marianozunino/jobby/dtos"
	"github.com/marianozunino/jobby/ent"
)

func UserToDto(entity *ent.User) *dtos.User {
	dto := &dtos.User{
		ID:       entity.ID,
		FullName: entity.FullName,
		Email:    entity.Email,
		Role:     dtos.Role(entity.Role),
	}
	return dto
}

func UsersToDtos(entities ent.Users) []*dtos.User {
	dtos := make([]*dtos.User, len(entities))
	for i, entity := range entities {
		dtos[i] = UserToDto(entity)
	}
	return dtos
}
