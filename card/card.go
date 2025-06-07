package card

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/eucatur/go-toolbox/text"
)

// Mask Hide card information
func Mask(cardNumber string) (cardMask string, err error) {
	isMasked := AlreadyMasked(cardNumber)

	if isMasked {
		return cardNumber, nil
	}

	if !Valid(cardNumber) {
		err = errors.New("invalid Card")
		return
	}

	regexMasked := regexp.MustCompile(`(?mi)(?P<SixFirst>\d{6})\d{3,9}(?P<LastFour>\d{4})`)

	useMask := strings.Repeat("X", len([]rune(cardNumber))-10)

	idxSixFirst := regexMasked.SubexpIndex("SixFirst")
	idxLastFour := regexMasked.SubexpIndex("LastFour")

	match := regexMasked.FindStringSubmatch(cardNumber)

	if len(match) <= 0 {
		l := len(cardNumber)

		cardMask = fmt.Sprintf("%s%s%s", cardNumber[0:6], strings.Repeat("X", l-10), cardNumber[l-4:l])
		return
	}

	cardMask = fmt.Sprintf("%s%s%s", match[idxSixFirst], useMask, match[idxLastFour])

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

	return l >= 13 && l <= 19
}

func GetInicialBin(cardNumber string) string {

	cardNumber = strings.Join(strings.Fields(strings.TrimSpace(cardNumber)), "")

	if len([]rune(cardNumber)) >= 6 {
		return cardNumber[:6]
	}

	return ""

}

func GetFinalBin(cardNumber string) string {

	cardNumber = strings.Join(strings.Fields(strings.TrimSpace(cardNumber)), "")

	if len([]rune(cardNumber)) >= 4 {
		return cardNumber[len([]rune(cardNumber))-4:]
	}

	return ""
}

func AlreadyMasked(cardNumber string) bool {

	if text.StringIsEmptyOrWhiteSpace(cardNumber) {
		return false
	}

	if len(cardNumber) < 13 {
		return false
	}

	regexMasked := regexp.MustCompile(`(?mi)(?P<SixFirst>\d{6})\d{3,9}(?P<LastFour>\d{4})`)

	idxSixFirst := regexMasked.SubexpIndex("SixFirst")
	idxLastFour := regexMasked.SubexpIndex("LastFour")

	match := regexMasked.FindStringSubmatch(cardNumber)

	if len(match) == 0 {
		return true
	}

	sixFirst := match[idxSixFirst]
	lastFour := match[idxLastFour]

	maskApplied := strings.ReplaceAll(cardNumber, sixFirst, "")
	maskApplied = strings.ReplaceAll(maskApplied, lastFour, "")

	hasMask := text.OnlyNumbers(maskApplied) == ""

	return strings.Index(cardNumber, sixFirst) == 1 &&
		strings.LastIndex(cardNumber, lastFour) != -1 &&
		hasMask

}
