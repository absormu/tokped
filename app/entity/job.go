package entity

type JobList struct {
	ID          string `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	Url         string `json:"url,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	Company     string `json:"company,omitempty"`
	CompanyUrl  string `json:"company_url,omitempty"`
	Location    string `json:"location,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	HowToApply  string `json:"how_to_apply,omitempty"`
	CompanyLogo string `json:"company_logo,omitempty"`
}
