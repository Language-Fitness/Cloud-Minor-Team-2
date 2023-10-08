package entity

type Module struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Difficulty  int      `json:"difficulty"`
	Tags        []string `json:"tags"`
	MadeBy      string   `json:"made_by"`
	Private     bool     `json:"private"`
	Key         string   `json:"key"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	SoftDeleted string   `json:"soft_deleted"`
}
