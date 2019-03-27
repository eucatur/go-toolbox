package duration

import (
    "fmt"
    "time"
)

// AsTextInBR formata uma duração (time.Duration) como
// texto em português (brasileiro).
func AsTextInBR(d time.Duration) string {
    text := ""
    hours := int64(d.Hours())
    mins := int64(d.Minutes()) % 60

    format := func(unit int64, desc string) string {
        str := ""
        if unit > 0 {
            str += fmt.Sprintf("%d %s", unit, desc)
            if unit > 1 {
                str += "s"
            }
        }
        return str
    }

    text += format(hours, "hora")

    if hours > 0 && mins > 0 {
        text += " e "
    }

    text += format(mins, "minuto")

    return text
}
