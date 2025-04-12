package validators

import (
	"fmt"
	"net/mail"
)

func Email(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("invalid email: %s", err.Error())
	}

	return nil
}
