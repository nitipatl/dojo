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

func Test_input_9_digitDisplay_9(t *testing.T) {
	output := " _ \n"
	output += "|_|\n"
	output += " _|"
	if digit(9) != output {
		t.Error("actual ", digit(9))
	}
}

func Test_input_29_digitDisplay_29(t *testing.T) {
	output := " _  _ \n"
	output += " _||_|\n"
	output += "|_  _|"
	if digit(29) != output {
		t.Error("actual ", digit(29))
	}
}

func Test_input_minus1_digitDisplay_minus1(t *testing.T) {
	output := "    \n"
	output += "-  |\n"
	output += "   |"
	if digit(-1) != output {
		t.Error("actual ", digit(-1))
	}
}
