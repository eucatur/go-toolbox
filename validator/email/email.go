package email

import "regexp"

// Valido verica se o valor informado é um email válido.
func Valido(email string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,6}$`).MatchString(email)
}
