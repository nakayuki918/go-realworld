package entity

type User struct {
	UserID string `json:"user_id"`
	Email string `json:"email"`
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}