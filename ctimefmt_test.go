package ctimefmt

import "time"
import "testing"

var format string = "%Y-%m-%d %H:%M:%S.%f"
var value string = "2019-01-02 03:04:05.666666"

func TestFormat(t *testing.T) {
	dt := time.Date(2019, 1, 2, 3, 4, 5, 666666000, time.UTC)
	s := Format(format, dt)
	if s != value {
		t.Errorf("Given: %v, expected: %v", s, value)
	}
}

func TestParse(t *testing.T) {
	dt := time.Date(2019, 1, 2, 3, 4, 5, 666666000, time.UTC)
	dt_, err := Parse(format, value)
	if err != nil {
		t.Error(err)
	} else if dt != dt {
		t.Errorf("Given: %v, expected: %v", dt_, dt)
	}
}
