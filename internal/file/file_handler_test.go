package file

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFileNames(t *testing.T) {
	actual, err := GetFileNames("testdata")
	if err != nil {
		t.Error(err)
	}

	expected := []string{"test.json"}
	assert.Equal(t, expected, actual)
}
