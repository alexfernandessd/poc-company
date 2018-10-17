package company

// Company is a model
type Company struct {
	ID      string `json:"id" structs:"id"`
	Company string `json:"company" structs:"company"`
}
