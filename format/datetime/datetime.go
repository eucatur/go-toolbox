package datetime

import (
	"log"
	"time"
)

func EUAtoBR(str string) string {

	if d, e := time.Parse("2006-01-02 15:04:05", str); e == nil {
		return d.Format("02/01/2006 15:04:05")
	}

	return ""
}

func EUAtoBRShort(str string) string {

	if d, e := time.Parse("2006-01-02 15:04:05", str); e == nil {
		return d.Format("02/01/2006 15:04")
	}

	return ""
}

func AsBRShort(dt time.Time) string {
	return dt.Format("02/01/2006 15:04")
}

const (
	DateTimeEUALayout = "2006-01-02 15:04:05"
	DateTimeBRLayout  = "02/01/2006 15:04:05"
)

func MustStrBRParseDateTimeBR(str string) time.Time {
	time, err := time.Parse(DateTimeBRLayout, str)
	if err != nil {
		log.Panic(err)
	}
	return time

}

func MustStrEUAParseDateTimeEUA(str string) time.Time {
	time, err := time.Parse(DateTimeEUALayout, str)
	if err != nil {
		log.Panic(err)
	}
	return time

}

func MustStrEUAParseStrBR(str string) string {
	time, err := time.Parse(DateTimeEUALayout, str)
	if err != nil {
		log.Panic(err)
	}
	return time.Format(DateTimeBRLayout)

}

func MustStrBRParseStrEUA(str string) string {
	time, err := time.Parse(DateTimeBRLayout, str)
	if err != nil {
		log.Panic(err)
	}
	return time.Format(DateTimeEUALayout)

}

func MustStrBRParseDateTimeEUA(str string) time.Time {
	return MustStrEUAParseDateTimeEUA(MustStrBRParseStrEUA(str))
}

func MustStrEUAParseDateTimeBR(str string) time.Time {
	return MustStrBRParseDateTimeBR(MustStrEUAParseStrBR(str))
}

func MustDateTimeEUAParseStrEUA(time time.Time) string {
	return time.Format(DateTimeEUALayout)
}

func MustDateTimeBRParseStrBR(time time.Time) string {
	return time.Format(DateTimeBRLayout)
}
