package duration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDurationISO8061(t *testing.T) {
	actual, err := ParseDurationISO8061("PT5.5S")
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, "5.5s", actual.String())
}
