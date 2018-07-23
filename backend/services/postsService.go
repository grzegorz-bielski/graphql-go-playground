package services

import "context"

// PostsService handles operations on posts
type PostsService struct{}

// NewPostsService is creating new PostService
func NewPostsService() *PostsService {
	return &PostsService{}
}

func (ps PostsService) GetPost(ctx context.Context, ID string) (Post, error) {
	return Post{}, nil
}
