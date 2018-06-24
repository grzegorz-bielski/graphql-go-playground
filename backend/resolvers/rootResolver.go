package resolvers

import "grzegorz-bielski/microstream/backend/services"

// RootResolver contains all schema resolvers
type RootResolver struct {
	*PostsResolver
}

// NewRootResolver returns new RootResolver
func NewRootResolver(rootService services.RootService) *RootResolver {
	return &RootResolver{
		NewPostsResolver(rootService),
	}
}
