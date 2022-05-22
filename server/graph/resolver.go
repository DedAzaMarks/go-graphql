package graph

import "github.com/DedAzaMarks/hw07-graphql/server/repository"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Repo repository.Repository
}

func NewResolver(repo repository.Repository) *Resolver {
	return &Resolver{Repo: repo}
}
