package graph

import (
	"github.com/marianozunino/jobby/graph/dataloader"
	"github.com/marianozunino/jobby/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	service.Service
	DataLoaders dataloader.Retriever
}
