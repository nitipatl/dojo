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
	output += digitNumber[input][0] + "\n"
	output += digitNumber[input][1] + "\n"
	output += digitNumber[input][2]
	return output
}
