package repo

import (
	// "database/sql"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `db:"id" json:"id"`
	FirstName   string `db:"first_name" json:"firstName"`
	LastName    string `db:"last_name" json:"lastName"`
	Email       string `db:"email" json:"email"`
	Password    string `db:"password" json:"password"`
	IsShopOwner bool   `db:"is_shop_owner" json:"isShopOwner"`
}

type UserRepo interface {
	Create(usr User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	dbCon *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	return &userRepo{
		dbCon: dbCon,
	}
}

func (r userRepo) Create(usr User) (*User, error) {
	query := `
INSERT INTO users (
    first_name,
    last_name,
    email,
    password,
    is_shop_owner
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id;
`

	var userID int
	err := r.dbCon.QueryRow(
		query,
		usr.FirstName,
		usr.LastName,
		usr.Email,
		usr.Password,
		usr.IsShopOwner,
	).Scan(&userID)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &usr, nil
}

// func (r userRepo) Find(email, pass string) (*User, error) {
// 	for _, u := range r.users {
// 		if u.Email == email && u.Password == pass {
// 			return &u, nil
// 		}
// 	}

// 	return nil, nil
// }

func (r userRepo) Find(email, pass string) (*User, error) {
	query := `
	SELECT
		id, first_name, last_name, email, password, is_shop_owner
	FROM users
	WHERE email = $1 AND password = $2
	LIMIT 1;
	`

	var user User
	err := r.dbCon.Get(&user, query, email, pass)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}
