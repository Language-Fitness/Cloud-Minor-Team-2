package entity

type Course struct {
	ID          int      `json:"id"`
	ModuleID    int      `json:"module_id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Difficulty  int      `json:"difficulty"`
	Tags        []string `json:"tags"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	SoftDeleted string   `json:"soft_deleted"`
}
