package database

import (
	"github.com/aadi-1024/auth-micro/pkg/jwtUtil"
	"github.com/aadi-1024/auth-micro/pkg/models"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

// VerifyLogin verifies whether the credentials are correct and if so returns a JWT token
func (d *Database) VerifyLogin(user models.User, j *jwtUtil.JwtConfig) (string, error) {
	query := `select uid, password from users where username = $1;`

	tx, err := d.Pool.BeginTx(ctx(), pgx.TxOptions{})
	if err != nil {
		return "", err
	}
	row := tx.QueryRow(ctx(), query, user.Username)
	var pass string
	var uid int
	err = row.Scan(&uid, &pass)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(user.Password))
	if err != nil {
		return "", err
	}
	//if no err up till this point, password has been verified as correct
	token, err := j.GenerateToken(uid)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (d *Database) RegisterUser(user models.User) error {
	query := `insert into users (username, password) values ($1, $2);`

	passHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), -1)
	if err != nil {
		return err
	}

	tx, err := d.Pool.BeginTx(ctx(), pgx.TxOptions{})
	defer tx.Rollback(ctx())
	if err != nil {
		return err
	}
	_, err = tx.Exec(ctx(), query, user.Username, string(passHash))
	if err != nil {
		return err
	}
	return tx.Commit(ctx())
}

func (d *Database) ResetPassword(user models.User, new_pass []byte) error {
	query_getuser := `select password from users where username = $1;`
	query_update := `update users set password = $1 where username = $2;`
	u := models.User{}

	tx, err := d.Pool.BeginTx(ctx(), pgx.TxOptions{})
	defer tx.Rollback(ctx())
	if err != nil {
		return err
	}

	row := tx.QueryRow(ctx(), query_getuser, user.Username)
	if err = row.Scan(&u.Password); err != nil {
		return err
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(user.Password)); err != nil {
		return err
	}
	newHash, err := bcrypt.GenerateFromPassword(new_pass, -1)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx(), query_update, string(newHash), user.Username)
	if err != nil {
		return err
	}
	return tx.Commit(ctx())
}
