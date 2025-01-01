package dto

type UserDto struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Nickname  string `json:"nickname"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
