package path

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// GetPackage return package path
// Example:
//   packagePath = GetPackage("/bitbucket.org/eucatur/")
func GetPackage(repository string) (packagePath string) {
	dir, _ := os.Getwd()
	path := strings.Split(dir, repository)
	p := strings.Split(path[1], "/")
	return path[0] + repository + p[0] + "/"
}

// GetRelativePath returns the path to the file that called this function without the project root path.
func GetRelativePath() (relativePath string, err error) {
	rootPath, err := os.Getwd()
	if err != nil {
		return
	}

	_, fullPathFile, _, ok := runtime.Caller(0)
	if !ok {
		err = errors.New("unable to identify file path")
		return
	}

	relativePathFile, err := filepath.Rel(rootPath, fullPathFile)
	if err != nil {
		return
	}

	relativePath = filepath.Dir(relativePathFile)
	return
}
