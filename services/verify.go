package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/utakatalp/email-verifier/repositories"
)

var (
	ErrExpired         = errors.New("activation link expired")
	ErrAlreadyVerified = errors.New("email already verified")
	ErrEmailExists     = errors.New("this email has already been registered")
	ErrTokenExists     = errors.New("token conflict, please try again")
	ErrNotFound        = errors.New("invalid token")
	ErrMailSendFailed  = errors.New("failed to send activation email")
)

func StartVerification(email string) (string, error) {
	token := uuid.New().String()
	expiresAt := time.Now().Add(10 * time.Minute)

	err := repositories.SaveVerification(email, token, expiresAt)
	if err != nil {
		return "", err
	}

	if err := SendActivationMail(email, token); err != nil {
		return "", err
	}

	return token, nil
}

func CompleteVerification(token string) (string, error) {
	v, err := repositories.GetVerificationByToken(token)
	if err != nil {
		return "", err
	}

	if time.Now().After(v.ExpiresAt) {
		return "", ErrExpired
	}

	if v.Verified {
		return "", ErrAlreadyVerified
	}

	if err := repositories.MarkVerified(token); err != nil {
		return "", err
	}

	return v.Email, nil
}
