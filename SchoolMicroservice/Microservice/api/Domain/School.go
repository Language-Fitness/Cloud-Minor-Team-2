package Domain

type School struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Location    string `json:"location"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	SoftDeleted bool   `json:"soft_deleted"`
}
