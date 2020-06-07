package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/HydroOSS/HydroAPI/graph/generated"
	"github.com/HydroOSS/HydroAPI/graph/model"
)

func (r *userResolver) Msgs(ctx context.Context, obj *model.User, guildID string) (*int, error) {
	res := 0
	if msg, ok := r.msgs[guildID]; ok {
		res = msg
	}
	return &res, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r, nil} }

type userResolver struct {
	*Resolver
	msgs map[string]int
}
