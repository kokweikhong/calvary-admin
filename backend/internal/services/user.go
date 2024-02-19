package services

import (
	"database/sql"

	"github.com/kokweikhong/calvary-admin/backend/internal/db"
	"github.com/kokweikhong/calvary-admin/backend/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetUsers() ([]*models.User, error)
	GetUser(id int) (*models.User, error)
	CreateUser(user *models.User) (int, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id int) error
}

type userService struct {
	db *sql.DB
}

func NewUserService() UserService {
	return &userService{
		db: db.GetPostgres(),
	}
}

func (s *userService) GetUsers() ([]*models.User, error) {
	query := s.getQuery(0)

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	users := []*models.User{}
	for rows.Next() {
		user := new(models.User)
		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.Department,
			&user.ProfileImage,
			&user.IsActive,
			&user.Position,
			&user.UpdatedAt,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *userService) GetUser(id int) (*models.User, error) {
	query := s.getQuery(id)

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	user := new(models.User)
	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.Department,
			&user.ProfileImage,
			&user.IsActive,
			&user.Position,
			&user.UpdatedAt,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (s *userService) CreateUser(user *models.User) (int, error) {
	query := `INSERT INTO users (
		username,
		email,
		password,
		role,
		department,
		profile_image,
		is_active,
		position
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8
	) RETURNING id`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return -1, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}

	user.Password = string(hashedPassword)

	err = stmt.QueryRow(
		user.Username,
		user.Email,
		user.Password,
		user.Role,
		user.Department,
		user.ProfileImage,
		user.IsActive,
		user.Position,
	).Scan(&user.ID)
	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (s *userService) UpdateUser(user *models.User) (*models.User, error) {
	query := `UPDATE users SET
		username = $1,
		email = $2,
		role = $3,
		department = $4,
		profile_image = $5,
		is_active = $6,
		position = $7,
		updated_at = now()
	WHERE id = $8`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		user.Username,
		user.Email,
		user.Role,
		user.Department,
		user.ProfileImage,
		user.IsActive,
		user.Position,
		user.ID,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(id int) error {
	query := `DELETE FROM users WHERE id = $1`

	stmt, err := s.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

// optional get query string either get all or get by id
func (s *userService) getQuery(id int) string {
	query := `SELECT
		id,
		username,
		email,
		password,
		role,
		department,
		profile_image,
		is_active,
		position,
		updated_at,
		created_at
	FROM users`

	if id > 0 {
		query += ` WHERE id = $1`
	}

	return query
}
