// log in as super set of the standart golang "log" with somes new methods
package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/fatih/color"
)

// Print in terminal the error message and the line of code with de error
func Error(e error) {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])

	hostname, err := os.Hostname()

	if err != nil {
		log.Println(err)
	}

	color.Set(color.FgHiRed)

	fmt.Printf("ERROR: %s | FILE: %s:%d | LINE: %d | FUNCTION: %s | HOSTNAME: %s\n", e.Error(), file, line, line, f.Name(), hostname)

	color.Unset()
}

// For use the same Println for the standart log lib
func Println(a ...interface{}) {
	log.Println(a)
}

// Save ou create a new log file with errors
func File(file, text string) error {
	path := "logs"

	_, err := os.Stat(path)

	if err != nil {

		err = os.Mkdir(path, os.FileMode(511))

		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile(path+"/"+file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		return err
	}

	f.WriteString(time.Now().Format("2006-02-01 15:04:05") + " | " + text + "\n")
	f.Close()
	return nil
}
