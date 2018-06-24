package services

// RootService holds all services
type RootService struct {
	posts *PostsService
}

// NewRootService is creating new RootService
func NewRootService() *RootService {
	return &RootService{
		posts: NewPostsService(),
	}
}
