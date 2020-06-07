package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/HydroOSS/HydroAPI/graph/generated"
	"github.com/HydroOSS/HydroAPI/graph/model"
	"gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func (r *queryResolver) User(ctx context.Context, id *string) (*model.User, error) {
	cursor, err := rethinkdb.DB("Hydro").Table("servers").Get(id).Run(r.DB)
	if err != nil {
		return nil, err
	}

	var user model.User
	cursor.One(&user)
	return &user, nil
}

func (r *queryResolver) Server(ctx context.Context, id *string) (*model.Server, error) {
	return &model.Server{}, nil
}

func (r *queryResolver) Giveaway(ctx context.Context, id *string) (*model.Giveaway, error) {
	return &model.Giveaway{}, nil
}

func (r *queryResolver) Strike(ctx context.Context, id *string) (*model.Strike, error) {
	return &model.Strike{}, nil
}

func (r *queryResolver) Merit(ctx context.Context, id *string) (*model.Merit, error) {
	return &model.Merit{}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
