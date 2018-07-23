package loaders

import (
	"context"
	"errors"
	"grzegorz-bielski/microstream/backend/services"
	"sync"

	"github.com/graph-gophers/dataloader"
)

func LoadPost(ctx context.Context, ID string) (services.Post, error) {
	var post services.Post

	ctxLoader, err := extract(ctx, postLoaderKey)
	if err != nil {
		return post, err
	}

	data, err := ctxLoader.Load(ctx, dataloader.StringKey(ID))()
	if err != nil {
		return post, err
	}

	post, ok := data.(services.Post)
	if !ok {
		return post, errors.New("Wrong type")
	}

	return post, nil
}

func LoadPosts(ctx context.Context, IDs []string) ([]services.Post, error) {
	var posts []services.Post
	return posts, nil
}

type postLoader struct {
	rootService services.RootService
}

func newPostLoader(rootService services.RootService) dataloader.BatchFunc {
	return postLoader{rootService}.loadBatch
}

func (pl postLoader) loadBatch(ctx context.Context, ids dataloader.Keys) []*dataloader.Result {
	var (
		n       = len(ids)
		results = make([]*dataloader.Result, n)
		wg      sync.WaitGroup
	)

	wg.Add(n)

	for i, id := range ids {
		go func(i int, url dataloader.Key) {
			defer wg.Done()

			resp, err := pl.rootService.GetPost(ctx, id.String())
			results[i] = &dataloader.Result{Data: resp, Error: err}
		}(i, id)
	}

	wg.Wait()

	return results
}
