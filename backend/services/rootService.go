package services

// RootService holds all services
type RootService struct {
	*PostsService
}

// NewRootService is creating new RootService
func NewRootService() *RootService {
	return &RootService{
		NewPostsService(),
	}
}
