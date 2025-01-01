package dto

type CreateGymDto struct {
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	CreatedBy string `json:"createdBy"`
}

type GymDto struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Nickname  string  `json:"nickname"`
	Active    bool    `json:"active"`
	CreatedBy string  `json:"createdBy"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
	DeletedAt *string `json:"deletedAt"`
}
