package weather

import "testing"

func Test_input_1_digitDisplay_1(t *testing.T) {
	output := "  "
	output += " |"
	output += " |"
	if digit(1) != output {
		t.Error("actual ", digit(1))
	}

}
