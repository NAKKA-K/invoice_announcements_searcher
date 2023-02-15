package duration

import (
	"time"

	"github.com/sosodev/duration"
)

func ParseDurationISO8061(meiliDuration string) (*time.Duration, error) {
	d, err := duration.Parse(meiliDuration)
	if err != nil {
		return nil, err
	}

	duration := d.ToTimeDuration()
	return &duration, nil
}
