package items

// Service interface
type Service interface {
}

type service struct {
}

// NewService start the new service
func NewService() (Service, error) {
	return &service{}, nil
}
