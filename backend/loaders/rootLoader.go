package loaders

import (
	"context"
	"grzegorz-bielski/microstream/backend/services"

	"github.com/graph-gophers/dataloader"
)

const (
	postLoaderKey contextKey = "post"
)

// RootLoader is reponsible for loading ids into loaders' batch functions
type RootLoader struct {
	batchFuncs map[contextKey]dataloader.BatchFunc
}

// NewRootLoader creates new RootLoader
func NewRootLoader(rootService services.RootService) *RootLoader {
	return &RootLoader{
		batchFuncs: map[contextKey]dataloader.BatchFunc{
			postLoaderKey: newPostLoader(rootService),
		},
	}
}

// Attach is adding loaders to current request's context
func (rl RootLoader) Attach(ctx context.Context) context.Context {
	for key, batchFunc := range rl.batchFuncs {
		ctx = context.WithValue(ctx, key, dataloader.NewBatchedLoader(batchFunc))
	}

	return ctx
}
