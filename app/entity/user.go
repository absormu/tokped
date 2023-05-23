package entity

type User struct {
	ID            int64  `json:"user_id,omitempty"`
	Name          string `json:"name,omitempty"`
	Email         string `json:"email,omitempty"`
	UserContactID int    `json:"user_contact_id,omitempty"`
	RoleID        int    `json:"role_id,omitempty"`
	Password      string `json:"password,omitempty"`
	IsDeleted     int    `json:"is_deleted,omitempty"`
	CreatedAt     string `json:"created_at,omitempty"`
	CreatedBy     string `json:"created_by,omitempty"`
	ModifiedAt    string `json:"modified_at,omitempty"`
	ModifiedBy    string `json:"modified_by,omitempty"`
}

type UserData struct {
	ID         int64  `json:"user_id,omitempty"`
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	Active     *bool  `json:"active"`
	IsDeleted  int    `json:"is_deleted,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	CreatedBy  string `json:"created_by,omitempty"`
	ModifiedAt string `json:"modified_at,omitempty"`
	ModifiedBy string `json:"modified_by,omitempty"`
}
