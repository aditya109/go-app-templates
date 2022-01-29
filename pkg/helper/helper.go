package helper

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetAbsolutePath provides absolute path for a relative path - 
func GetAbsolutePath(relPath string) string {
	if relPath[0] == '/' {
		relPath = relPath[1:]
	}
	splitRelativePath := strings.Split(relPath, "/")
	cwd, err := os.Getwd()
	var path string
	if err != nil {
		log.Fatal(err)
	}
	projectLocation := strings.Split(cwd, "go-server-template")
	path = filepath.Join(projectLocation[0], "go-server-template", splitRelativePath[0], splitRelativePath[1])
	return path
}
