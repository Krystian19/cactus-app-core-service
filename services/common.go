package services

type server struct{}

// NewServer : Returns a pointer to a new Server
func NewServer() *server {
	return &server{}
}
