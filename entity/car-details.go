package entity

type CarDetails struct {
	ID             int    `json:"id"`
	Brand          string `json:"brand"`
	Model          string `json:"model"`
	Year           int    `json:"model_year"`
	Vin            string `json:"vin"`
	OwnerFirstName string `json:"owner_firstname"`
	OwnerLastName  string `json:"owner_lastname"`
	OwnerEmail     string `json:"owner_email"`
	OwnerJobTitle  string `json:"owner_job_title"`
}
