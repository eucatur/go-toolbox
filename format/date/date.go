package date

import (
	"log"
	"time"
)

const (
	DateEUALayout = "2006-01-02"
	DateBRLayout  = "02/01/2006"
)

func EUAParseBR(str string) string {

	if d, e := time.Parse(DateEUALayout, str); e == nil {
		return d.Format(DateBRLayout)
	}

	return ""
}

func AsEUA(dt time.Time) string {
	return dt.Format(DateEUALayout)
}

func MustStrBRParseDateBR(str string) time.Time {
	time, err := time.Parse(DateBRLayout, str)
	if err != nil {
		log.Panic(err)
	}
	return time

}

func MustStrEUAParseDateEUA(str string) time.Time {
	time, err := time.Parse(DateEUALayout, str)
	if err != nil {
		log.Panic(err)
	}
	return time

}

func MustStrEUAParseStrBR(str string) string {
	time, err := time.Parse(DateEUALayout, str)
	if err != nil {
		log.Panic(err)
	}
	return time.Format(DateBRLayout)

}

func MustStrBRParseStrEUA(str string) string {
	time, err := time.Parse(DateBRLayout, str)
	if err != nil {
		log.Panic(err)
	}
	return time.Format(DateEUALayout)

}

func MustStrBRParseDateEUA(str string) time.Time {
	return MustStrEUAParseDateEUA(MustStrBRParseStrEUA(str))
}

func MustStrEUAParseDateBR(str string) time.Time {
	return MustStrBRParseDateBR(MustStrEUAParseStrBR(str))
}

func MustDateEUAParseStrEUA(time time.Time) string {
	return time.Format(DateEUALayout)
}

func MustDateBRParseStrBR(time time.Time) string {
	return time.Format(DateBRLayout)
}
