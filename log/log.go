// Package log in as super set of the standart golang "log" with somes new methods
package log

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/fatih/color"
)

var debug = false

// EnableDebug active the debug mode
func EnableDebug(enable bool) {
	debug = enable
}

// Error print in terminal the error message and the line of code with de error
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

// Println use the same Println for the standart log lib
func Println(i ...interface{}) {
	log.Println(i...)
}

// Fatal use the same Fatal for the standart log lib
func Fatal(i ...interface{}) {
	log.Fatal(i...)
}

// Debug is a function for print in terminal if the variable debug it's true
func Debug(a ...interface{}) {
	if debug {
		log.Println(a...)
	}
}

// Debugf is a function for print formated in terminal if the variable debug it's true
func Debugf(format string, v ...interface{}) {
	if debug {
		log.Printf(format, v...)
	}
}

// File save ou create a new log file with errors
func File(file, text string) error {
	if len(file) < 1 {
		return errors.New("invalid file name")
	}

	path := "logs/"
	if file[0:1] == "/" {
		path += file[1:]
	} else {
		path += file
	}

	pathSplited := strings.Split(path, "/")
	folders := strings.Join(pathSplited[0:len(pathSplited)-1], "/")

	_, err := os.Stat(folders)

	if err != nil {

		err = os.MkdirAll(folders, os.ModePerm)

		if err != nil {
			return err
		}
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)

	if err != nil {
		return err
	}

	f.WriteString(time.Now().Format("2006-01-02 15:04:05") + " | " + text + "\n")
	f.Close()
	return nil
}

func createFileIfNotExists(filePath string) (f *os.File) {
	dirName := filepath.Dir(filePath)

	err := os.MkdirAll(dirName, os.ModePerm)

	if err != nil {
		log.Fatalf("Error create path: %v", err)
	}

	f, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalf("Error create file: %v", err)
	}

	return
}

//SetOutputFiles set a new path for logs
func SetOutputFiles(outfilePath, errfilePath string) {
	outFile := createFileIfNotExists(outfilePath)
	errFile := createFileIfNotExists(errfilePath)

	defer outFile.Close()
	defer errFile.Close()

	syscall.Dup2(int(outFile.Fd()), 1) /* -- stdout */
	syscall.Dup2(int(errFile.Fd()), 2) /* -- stderr */
}

// ReturnDetailsError return the error message and the line of code with de error
func ReturnDetailsError(e error) (buildedErrorDetails string) {
	pc := make([]uintptr, 10)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])

	hostname, err := os.Hostname()

	if err != nil {
		log.Println(err)
	}

	buildedErrorDetails = fmt.Sprintf("ERROR: %s | FILE: %s:%d | LINE: %d | FUNCTION: %s | HOSTNAME: %s\n", e.Error(), file, line, line, f.Name(), hostname)

	return
}
