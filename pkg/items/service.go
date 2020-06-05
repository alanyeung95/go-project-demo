package items

// Service interface
type Service interface {
}

type service struct {
	repository Repository
}

// NewService start the new service
func NewService(repository Repository) (Service, error) {
	return &service{repository}, nil
}
