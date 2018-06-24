package loaders

import (
	"context"
	"fmt"

	"github.com/graph-gophers/dataloader"
)

type contextKey string

func (ck contextKey) String() string {
	return string(ck)
}

func extract(ctx context.Context, key contextKey) (*dataloader.Loader, error) {
	ldr, ok := ctx.Value(key).(*dataloader.Loader)
	if !ok {
		return nil, fmt.Errorf("unable to find %s loader on the request context", key)
	}

	return ldr, nil
}
