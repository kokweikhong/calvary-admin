package models

type User struct {
	ID           int    `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	Email        string `json:"email" db:"email"`
	Password     string `json:"password" db:"password"`
	Role         string `json:"role" db:"role"`
	Department   string `json:"department" db:"department"`
	ProfileImage string `json:"profile_image" db:"profile_image"`
	IsActive     bool   `json:"is_active" db:"is_active"`
	Position     string `json:"position" db:"position"`
	UpdatedAt    string `json:"updated_at" db:"updated_at"`
	CreatedAt    string `json:"created_at" db:"created_at"`
}
