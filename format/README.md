# Format #
É um lib com funções de formatação para diversos tipos

## date ##

func EUAtoBR(str string) string {}

func AsEUA(dt time.Time) string {}

## datetime ##

func EUAtoBR(str string) string {}

func EUAtoBRShort(str string) string {}

func AsBRShort(dt time.Time) string {}

## duration ##

func AsTextInBR(d time.Duration) string {}

## money ##

func Reais(valor int64) string {}

func Round(value float64, precision int) float64 {}

## phone ##

func Short(phone string) string {}

## time ##

func HourMin(d time.Duration) string {}

func AsDefault(dt time.Time) string {}

func ToShort(dt string) string {}
