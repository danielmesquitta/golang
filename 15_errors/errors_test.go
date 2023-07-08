package errorsPkg

import "testing"

func TestParseTime(t *testing.T) {
	table := []struct {
		input string
		ok    bool
	}{
		{"14:07:33", true},
		{"1:02:30", true},
		{"XX:XX:XX", false},
		{"14:07", false},
		{"abc", false},
	}

	for _, data := range table {
		_, err := ParseTime(data.input)

		if data.ok && err != nil {
			t.Errorf("ParseTime(%s) should not return an error", data.input)
		}

		if !data.ok && err == nil {
			t.Errorf("ParseTime(%s) should return an error", data.input)
		}
	}
}
