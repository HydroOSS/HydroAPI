package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/HydroOSS/HydroAPI/graph/generated"
	"github.com/HydroOSS/HydroAPI/graph/model"
)

func (r *userResolver) Msgs(ctx context.Context, obj *model.User, guildID string) (*int, error) {
	res := 0
	return &res, errors.New("to be implemented")
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
