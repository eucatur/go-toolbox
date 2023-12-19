package commandredis

import (
	"database/sql/driver"
	"encoding/xml"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	gotoolboxtext "github.com/eucatur/go-toolbox/text"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CommandRedis int64

type NameType CommandRedis

const (
	Undefined CommandRedis = 0
	Delete    CommandRedis = 1 << iota
	Expire
	Get
	Match
	Ping
	Scan
	Set

	name_undefined = "undefined"
	name_del       = "DEL"
	name_expire    = "EXPIRE"
	name_get       = "GET"
	name_match     = "MATCH"
	name_ping      = "PING"
	name_scan      = "SCAN"
	name_set       = "SET"
)

var (
	commandredis_name = map[CommandRedis]string{
		Undefined: name_undefined,
		Delete:    name_del,
		Expire:    name_expire,
		Get:       name_get,
		Match:     name_match,
		Ping:      name_ping,
		Scan:      name_scan,
		Set:       name_set,
	}

	commandredis_value = map[string]CommandRedis{
		name_undefined: Undefined,
		name_del:       Delete,
		name_expire:    Expire,
		name_get:       Get,
		name_match:     Match,
		name_ping:      Ping,
		name_scan:      Scan,
		name_set:       Set,
	}
)

func TryParseToEnum(value string) CommandRedis {

	tpEnum := new(CommandRedis)

	err := tpEnum.Scan(value)

	if err != nil {
		return Undefined
	}

	return *tpEnum

}

func (c CommandRedis) OperationsStatemented() (operationsStatemented []CommandRedis) {

	if c == Undefined {

		return []CommandRedis{}

	}

	for _, tpEnum := range commandredis_value {

		if c&tpEnum != 0 {

			operationsStatemented = append(operationsStatemented, tpEnum)

		}

	}

	if len(operationsStatemented) > 1 {

		sort.Slice(operationsStatemented, func(i, j int) bool {

			return operationsStatemented[i] < operationsStatemented[j]

		})

		return

	}

	return []CommandRedis{}

}

func List() []CommandRedis {

	enu := []CommandRedis{}

	for enum := range commandredis_name {

		enu = append(enu, enum)

	}

	return enu

}

func (c CommandRedis) List() []CommandRedis {

	return List()

}

func (c CommandRedis) String() string {

	tpEnum, ok := commandredis_name[c]

	if !ok {

		return name_undefined

	}

	return tpEnum

}

func TypesAccepts() string {

	tpEnumAvailable := []CommandRedis{}

	for tpEnu := range commandredis_name {

		if tpEnu == Undefined {

			continue

		}

		tpEnumAvailable = append(tpEnumAvailable, tpEnu)

	}

	strTypeAvailables := ""

	for idx, tpEnum := range tpEnumAvailable {

		switch {

		case len(tpEnumAvailable)-1 == idx:

			strTypeAvailables += fmt.Sprintf(" ou %d|%s", tpEnum, tpEnum.String())

		case len(tpEnumAvailable)-2 == idx:

			strTypeAvailables += fmt.Sprintf("%d|%s", tpEnum, tpEnum.String())

		default:

			strTypeAvailables += fmt.Sprintf("%d|%s, ", tpEnum, tpEnum.String())

		}

	}

	return strTypeAvailables

}

func (c CommandRedis) TypesAccepts() string {

	return TypesAccepts()

}

func (c CommandRedis) IsValid() (err error) {

	check := c

	err = check.Scan(c)

	return
}

func (c CommandRedis) MarshalJSON() ([]byte, error) {

	return []byte(fmt.Sprintf(`"%s"`, c.String())), nil

}

func (c *CommandRedis) UnmarshalJSON(bytes []byte) error {

	value, err := c.tryGetValueFromJSON(bytes)

	if err == nil && !strings.EqualFold(value, "") {

		tpEnum, err := c.tryParseValueToTypeEnum(value)

		if err != nil {

			*c = Undefined

			return err

		}

		*c = CommandRedis(tpEnum)

	}

	return err

}

func ParseType(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {

	if f.Kind() != reflect.String || t != reflect.TypeOf(Undefined) {

		return data, nil

	}

	tpEnum := Undefined

	err := tpEnum.Scan(data)

	return tpEnum, err

}

func (c CommandRedis) Value() (driver.Value, error) {

	return c.String(), nil

}

func (c *CommandRedis) Scan(value interface{}) (err error) {

	defer func() {

		if errRecover := recover(); errRecover != nil {

			err = fmt.Errorf("Scan failed for value %v for type of CommandRedis. Details: %v", value, errRecover)

			*c = Undefined

		}

	}()

	switch data := value.(type) {

	case []uint8:

		if len(data) <= 0 {

			*c = Undefined

			return nil

		}

		str := string([]byte(data))

		tpEnum, err := c.tryParseValueToTypeEnum(str)

		if err != nil {

			*c = Undefined

			return err

		}

		*c = CommandRedis(tpEnum)

	case int:

		tpEnum, err := c.tryParseValueToTypeEnum(fmt.Sprintf("%d", data))

		if err != nil {

			*c = Undefined

			return err

		}

		*c = tpEnum

	case int32:

		tpEnum, err := c.tryParseValueToTypeEnum(fmt.Sprintf("%d", data))

		if err != nil {

			*c = Undefined

			return err

		}

		*c = tpEnum

	case float32:

		tpEnum, err := c.tryParseValueToTypeEnum(fmt.Sprintf("%d", int64(data)))

		if err != nil {

			*c = Undefined

			return err

		}

		*c = tpEnum

	case float64:

		tpEnum, err := c.tryParseValueToTypeEnum(fmt.Sprintf("%d", int64(data)))

		if err != nil {

			*c = Undefined

			return err

		}

		*c = tpEnum

	case int64:

		tpEnum, err := c.tryParseValueToTypeEnum(fmt.Sprintf("%d", data))

		if err != nil {

			*c = Undefined

			return err

		}

		*c = tpEnum

	case string:

		dataNoQuote, _ := strconv.Unquote(data)

		if !strings.EqualFold(dataNoQuote, "") {

			data = dataNoQuote

		}

		if strings.EqualFold(dataNoQuote, "") && (strings.EqualFold(data, "\"\"") || strings.EqualFold(data, "")) {

			*c = Undefined

			return err

		}

		tpEnum, err := c.tryParseValueToTypeEnum(data)

		if err != nil {

			*c = Undefined

			return err

		}

		*c = tpEnum

	case CommandRedis:

		tpEnum, err := c.tryParseValueToTypeEnum(fmt.Sprintf("%d", data))

		if err != nil {

			*c = Undefined

			return err

		}

		*c = tpEnum

	}

	return nil

}

func (c CommandRedis) MarshalXML(xmlEnc *xml.Encoder, start xml.StartElement) error {

	return xmlEnc.EncodeElement(c.String(), start)

}

func (c *CommandRedis) UnmarshalXML(xmlDec *xml.Decoder, start xml.StartElement) error {

	var (
		valueContent string
		tpEnum       CommandRedis
	)

	xmlDec.DecodeElement(&valueContent, &start)

	tpEnum, err := c.tryParseValueToTypeEnum(valueContent)

	if err != nil {

		return err

	}

	*c = CommandRedis(tpEnum)

	return err

}

func (c *CommandRedis) tryParseValueToTypeEnum(value string) (tpEnum CommandRedis, err error) {

	tpEnumINT, ok := c.tryGetEnumINT(value)

	if !ok {

		valueINT, err := strconv.Atoi(value)

		if err != nil {

			return Undefined, fmt.Errorf("the %s is incorret to type of CommandRedis", value)

		}

		tpEnumStr, ok := commandredis_name[CommandRedis(valueINT)]

		if !ok {

			return Undefined, fmt.Errorf("the %s not valid type of CommandRedis", value)

		}

		tpEnumINT, ok = commandredis_value[tpEnumStr]

		if !ok {

			return Undefined, fmt.Errorf("the %s is invalid type of CommandRedis", value)

		}

	}

	return CommandRedis(tpEnumINT), nil

}

func (c *CommandRedis) tryGetValueFromJSON(bytes []byte) (value string, err error) {

	value, err = strconv.Unquote(string(bytes))

	if err != nil {

		valueINT := int(Undefined)

		valueINT, err = strconv.Atoi(string(bytes))

		if err != nil {

			return string(bytes), nil

		}

		value = fmt.Sprintf("%d", valueINT)

		err = nil

	}

	return

}

func (c *CommandRedis) tryGetEnumINT(value string) (CommandRedis, bool) {

	modeWrite := []string{
		"upper",
		"title",
		"lower",
	}

	convertString := func(mode, value string) string {

		gotoolboxtext.RemoveAccents(value)

		switch mode {

		case "lower":

			return strings.ToLower(value)

		case "upper":

			return strings.ToUpper(value)

		case "title":

			c := cases.Title(language.BrazilianPortuguese)

			return c.String(value)

		default:
			return value
		}

	}

	value = strings.ToLower(value)

	tpEnumINT, ok := commandredis_value[value]

	if !ok {

		for _, mode := range modeWrite {

			for nameEnum, tpEnum := range commandredis_value {

				nameEnumNormalized := convertString(mode, nameEnum)

				valueNormalized := convertString(mode, value)

				if strings.EqualFold(nameEnumNormalized, valueNormalized) {

					tpEnumINT = tpEnum

					ok = true

					break

				}

			}

		}

	}

	return tpEnumINT, ok

}
