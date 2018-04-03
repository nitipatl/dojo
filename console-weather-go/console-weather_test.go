package weather

import "testing"

func Test_input_1_digitDisplay_1(t *testing.T) {
	output := "   \n"
	output += "  |\n"
	output += "  |"
	if digit(1) != output {
		t.Error("actual ", digit(1))
	}
}

func Test_input_2_digitDisplay_2(t *testing.T) {
	output := " _ \n"
	output += " _|\n"
	output += "|_ "
	if digit(2) != output {
		t.Error("actual ", digit(2))
	}
}
