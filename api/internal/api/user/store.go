package user

import "database/sql"

func (uh *UserHandler) CreateUser(u *User) error {
	query := "INSERT INTO users (email, password_hash, role, created_at) VALUES (?, ?, ?, ?)"
	result, err := uh.app.DB().Exec(query, u.Email, u.PasswordHash, u.Role, u.CreatedAt)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}

func (uh *UserHandler) EmailExists(email string) (bool, error) {
	var id int64
	query := "SELECT id FROM users WHERE email = ?"
	err := uh.app.DB().Get(&id, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return id > 0, nil
}

func (uh *UserHandler) GetUserByEmail(email string) (*User, error) {
	query := "SELECT id, email, password_hash, role, created_at FROM users WHERE email = ?"
	var user User
	err := uh.app.DB().Get(&user, query, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
