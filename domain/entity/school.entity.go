package entity

type School struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	SoftDeleted string `json:"soft_deleted"`
}
