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
		for i := 0; i < len(out); i++ {
			out[i] += digitNumber[number][i]
		}
	}
	output := out[0] + "\n" + out[1] + "\n" + out[2]
	return output
}
