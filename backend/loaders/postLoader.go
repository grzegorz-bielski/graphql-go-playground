package loaders

import (
	"context"
	"grzegorz-bielski/microstream/backend/services"

	"github.com/graph-gophers/dataloader"
)

type postLoader struct {
	rootService services.RootService
}

func newPostLoader(rootService services.RootService) dataloader.BatchFunc {
	return postLoader{rootService}.loadBatch
}

func (pl postLoader) loadBatch(ctx context.Context, urls dataloader.Keys) []*dataloader.Result {
	// todo
}
