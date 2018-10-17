package company

const (
	companyTable = "company"
)

// Repository map methods
type Repository interface {
	get(companyID string) (Company, error)
}

// RepositoryImpl map requiments interfaces
type RepositoryImpl struct {
	db Database
}

// NewRepository create a new repository
func NewRepository(db Database) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

func (r RepositoryImpl) get(companyID string) (Company, error) {
	var company Company
	err := r.db.get(companyTable, companyID, &company)
	return company, err
}
