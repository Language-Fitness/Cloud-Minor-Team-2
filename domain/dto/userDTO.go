package dto

type UserDTO struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	SchoolID        string `json:"school_id"`
	WhitelistModule []int  `json:"whitelist_module"`
	RoleID          string `json:"role_id"`
	Settings        string `json:"settings"`
}
