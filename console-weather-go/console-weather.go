package weather

import (
	"strconv"
)

func digit(input int) string {

	digitNumber := [][3]string{
		{
			" _ ",
			"| |",
			"|_|",
		},
		{
			"   ",
			"  |",
			"  |",
		},
		{
			" _ ",
			" _|",
			"|_ ",
		},
		{
			" _ ",
			" _|",
			" _|",
		},
		{
			"   ",
			"|_|",
			"  |",
		},
		{
			" _ ",
			"|_ ",
			" _|",
		},
		{
			" _ ",
			"|_ ",
			"|_|",
		},
		{
			" _ ",
			"  |",
			"  |",
		},
		{
			" _ ",
			"|_|",
			"|_|",
		},
		{
			" _ ",
			"|_|",
			" _|",
		},
	}

	out := [3]string{
		"",
		"",
		"",
	}
	for _, num := range strconv.Itoa(input) {
		number, _ := strconv.Atoi(string(num))
		out[0] += digitNumber[number][0]
		out[1] += digitNumber[number][1]
		out[2] += digitNumber[number][2]
	}
	output := out[0] + "\n" + out[1] + "\n" + out[2]
	return output
}
