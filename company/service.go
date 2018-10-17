package company

// Service map requirements
type Service struct {
	coRepository Repository
}

// NewService returns a service layer
func NewService(coRepository Repository) *Service {
	return &Service{coRepository: coRepository}
}

// Get returns a Compnay by id
func (s Service) Get(companyID string) (Company, error) {
	return s.coRepository.get(companyID)
}
