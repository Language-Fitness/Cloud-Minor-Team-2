package response

type UserResponseDTO struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	SchoolID        string `json:"school_id"`
	WhitelistModule []int  `json:"whitelist_module"`
	RoleID          string `json:"role_id"`
	Rating          string `json:"ratings"`
	Settings        string `json:"settings"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	SoftDeleted     string `json:"soft_deleted"`
}
