package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/DedAzaMarks/hw07-graphql/server/graph/converter"
	"github.com/DedAzaMarks/hw07-graphql/server/graph/generated"
	"github.com/DedAzaMarks/hw07-graphql/server/graph/model"
	"github.com/google/uuid"
)

func (r *mutationResolver) AddComment(ctx context.Context, gameID string, text string) (string, error) {
	parsed, err := uuid.Parse(gameID)
	if err != nil {
		return "", err
	}
	for _, game := range r.Repo.Games {
		if game.Id == parsed {
			game.Comments = append(game.Comments, text)
			return text, nil
		}
	}
	return "", fmt.Errorf("game not found, comment not added: %s", gameID)
}

func (r *queryResolver) Game(ctx context.Context, id string) (*model.Game, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	for _, game := range r.Repo.Games {
		if game.Id == parsed {
			return converter.GameToGame(game), nil
		}
	}
	return nil, fmt.Errorf("game not found: %s", id)
}

func (r *queryResolver) Games(ctx context.Context, ids []string) ([]*model.Game, error) {
	var res []*model.Game
	for _, id := range ids {
		parsed, err := uuid.Parse(id)
		if err != nil {
			return nil, err
		}
		for _, game := range r.Repo.Games {
			if game.Id == parsed {
				res = append(res, converter.GameToGame(game))
			}
		}
	}
	return res, nil
}

func (r *queryResolver) Player(ctx context.Context, username string) (*model.Player, error) {
	for _, player := range r.Repo.Players {
		if player.Username == username {
			return converter.PlayerToPlayer(player), nil
		}
	}
	return nil, fmt.Errorf("player not found: %s", username)
}

func (r *queryResolver) Filter(ctx context.Context, alive bool) ([]*model.Player, error) {
	var res []*model.Player
	for _, player := range r.Repo.Players {
		if player.IsAlive == alive {
			res = append(res, converter.PlayerToPlayer(player))
		}
	}
	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
