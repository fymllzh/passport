package user

import (
	"database/sql"
	"github.com/wuzehv/passport/service/db"
)

type User struct {
	Id       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func FindByEmail(email string) (User, error) {
	s := "select * from user where email =?"

	var u User
	err := db.Db.QueryRow(s, email).Scan(&u.Id, &u.Email, &u.Password)

	switch {
	case err == sql.ErrNoRows:
		return User{}, nil
	case err != nil:
		return User{}, err
	default:
		return u, nil
	}
}
