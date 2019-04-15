package card

import (
	"errors"
	"fmt"
	"strings"

	"github.com/eucatur/go-toolbox/text"
)

// Mask Hide card information
func Mask(cardNumber string) (cardMask string, err error) {
	if !Valid(cardNumber) {
		err = errors.New("Invalid Card")
		return
	}

	l := len(cardNumber)

	cardMask = fmt.Sprintf("%s%s%s", cardNumber[0:4], strings.Repeat("*", l-8), cardNumber[l-4:l])

	return
}

func Valid(cardNumber string) bool {
	cardNumber = text.OnlyNumbers(cardNumber)
	l := len(cardNumber)

	return l > 14 && l < 19
}