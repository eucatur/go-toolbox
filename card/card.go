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

	cardMask = fmt.Sprintf("%s%s%s", cardNumber[0:6], strings.Repeat("X", l-10), cardNumber[l-4:l])

	return
}

func TryMask(cardNumber string) (cardMask string) {
	cardMask, err := Mask(cardNumber)
	if err != nil {
		return cardNumber
	}

	return
}

// Valid method to validate credit card
func Valid(cardNumber string) bool {
	cardNumber = text.OnlyNumbers(cardNumber)
	l := len(cardNumber)

	return l > 13 && l < 19
}
