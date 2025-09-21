package repositories

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/utakatalp/email-verifier/db"
)

type Verification struct {
	Email     string
	Token     string
	ExpiresAt time.Time
	Verified  bool
}

var (
	ErrEmailExists = errors.New("email already exists")
	ErrTokenExists = errors.New("token already exists")
	ErrNotFound    = errors.New("verification not found")
)

func SaveVerification(email, token string, expiresAt time.Time) error {
	_, err := db.Pool.Exec(context.Background(),
		"INSERT INTO email_verifications (email, token, expires_at) VALUES ($1, $2, $3)",
		email, token, expiresAt,
	)
	if err != nil {
		if strings.Contains(err.Error(), "email_verifications_email_key") {
			return ErrEmailExists
		}
		if strings.Contains(err.Error(), "email_verifications_token_key") {
			return ErrTokenExists
		}
		return err
	}
	return err
}

func GetVerificationByToken(token string) (Verification, error) {
	var v Verification
	err := db.Pool.QueryRow(context.Background(),
		"SELECT email, token, expires_at, verified FROM email_verifications WHERE token=$1",
		token,
	).Scan(&v.Email, &v.Token, &v.ExpiresAt, &v.Verified)
	return v, err
}

func MarkVerified(token string) error {
	_, err := db.Pool.Exec(context.Background(),
		"UPDATE email_verifications SET verified=TRUE WHERE token=$1", token,
	)
	return err
}
