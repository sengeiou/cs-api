package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (r *mutationResolver) Upload(ctx context.Context, file graphql.Upload) (string, error) {
	return r.storage.Upload(ctx, file.File, file.Filename)
}
