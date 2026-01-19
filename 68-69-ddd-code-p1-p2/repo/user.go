package repo

import (
	// "database/sql"
	"database/sql"
	"ecom/domain"
	"ecom/user"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	dbCon *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	return &userRepo{
		dbCon: dbCon,
	}
}

func (r userRepo) Create(usr domain.User) (*domain.User, error) {
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

func (r userRepo) Find(email, pass string) (*domain.User, error) {
	query := `
	SELECT
		id, first_name, last_name, email, password, is_shop_owner
	FROM users
	WHERE email = $1 AND password = $2
	LIMIT 1;
	`

	var user domain.User
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
