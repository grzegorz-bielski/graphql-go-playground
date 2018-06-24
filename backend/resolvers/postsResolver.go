package resolvers

import "grzegorz-bielski/microstream/backend/services"

// PostsResolver resolves posts
type PostsResolver struct {
	rootService services.RootService
}

// NewPostsResolver returns new PostsResolver
func NewPostsResolver(rootService services.RootService) *PostsResolver {
	return &PostsResolver{rootService: rootService}
}
