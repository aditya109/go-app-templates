package helper

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGetAbsolutePathForCorrectRelativePath(t *testing.T) {
	result := GetAbsolutePath(`/config/config.yaml`)
	var expectedResult string
	cwd, err := os.Getwd()
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	expectedResult = filepath.Join(strings.Split(cwd, "pkg/helper")[0], "config/config.yaml")
	if result != expectedResult {
		t.Errorf("got %q, wanted %q", result, expectedResult)
	}
}
