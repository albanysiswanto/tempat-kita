package authModel

import (
	"tempat-kita/config"
	"tempat-kita/entities"

	"golang.org/x/crypto/bcrypt"
)

func Register(user entities.User) bool {
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	result, err := config.DB.Exec(`
	INSERT INTO user (email, name, password, role, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?)
	`, user.Email, user.Name, hashPassword, 1, user.Created_At, user.Updated_At)
	if err != nil {
		panic(err)
	}

	LastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return LastInsertId > 0
}

func GetEmailLogin(email string) (entities.User, error) {
	var user entities.User
	err := config.DB.QueryRow(`
	SELECT id, email, name, password, role, created_at, updated_at
	FROM user
	WHERE email = ?
	`, email).Scan(&user.Id, &user.Email, &user.Name, &user.Password, &user.Role, &user.Created_At, &user.Updated_At)
	if err != nil {
		panic(err)
	}
	return user, err
}

func VerifyPassword(hashPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
}
