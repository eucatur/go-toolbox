package path

import (
	"os"
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
