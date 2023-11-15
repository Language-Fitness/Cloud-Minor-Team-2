// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type School struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Location    *string `json:"location,omitempty"`
	CreatedAt   *string `json:"created_at,omitempty"`
	UpdatedAt   *string `json:"updated_at,omitempty"`
	SoftDeleted *bool   `json:"soft_deleted,omitempty"`
}

type SchoolInput struct {
	Name     string  `json:"name"`
	Location *string `json:"location,omitempty"`
}
