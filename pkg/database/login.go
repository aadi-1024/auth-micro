package database

import (
	"github.com/aadi-1024/auth-micro/pkg/jwtUtil"
	"github.com/aadi-1024/auth-micro/pkg/models"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

// VerifyLogin verifies whether the credentials are correct and if so returns a JWT token
func (d *Database) VerifyLogin(user models.User, j *jwtUtil.JwtConfig) (string, error) {
	query := `select uid, password from users where username = $1`

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
