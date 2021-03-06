package models

import "errors"

var ErrInvalidEmailCode = errors.New("Invalid or expired email code")

type SendEmailCommand struct {
	To       []string
	Template string
	Data     map[string]interface{}
	Massive  bool
	Info     string
}
