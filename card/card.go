package card

import (
	"errors"
	"fmt"
	"strings"

	"github.com/eucatur/go-toolbox/text"
)

// Mask Hide card information
func Mask(cardNumber string) (cardMask string, err error) {
	cardNumber = text.OnlyNumbers(cardNumber)
	if !Valid(cardNumber) {
		err = errors.New("Invalid Card")
		return
	}

	l := len(cardNumber)

	cardMask = fmt.Sprintf("%s%s%s", cardNumber[0:4], strings.Repeat("*", l-8), cardNumber[l-4:l])

	return
}

func Valid(cardNumber string) bool {
	l := len(cardNumber)

	return l > 15 && l < 19
}
