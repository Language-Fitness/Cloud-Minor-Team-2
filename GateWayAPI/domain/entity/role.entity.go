package entity

type Role struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Permissions []string `json:"permissions"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	SoftDeleted string   `json:"soft_deleted"`
}
