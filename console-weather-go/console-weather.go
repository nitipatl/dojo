package weather

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

	output := ""
	if input == 29 {
		output += digitNumber[2][0] + digitNumber[9][0] + "\n"
		output += digitNumber[2][1] + digitNumber[9][1] + "\n"
		output += digitNumber[2][2] + digitNumber[9][2]
	} else {
		output += digitNumber[input][0] + "\n"
		output += digitNumber[input][1] + "\n"
		output += digitNumber[input][2]
	}

	return output
}
