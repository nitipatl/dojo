package weather

func digit(input int) string {

	top := []string{"   ", "   "}
	mid := []string{"   ", "  |"}
	bot := []string{"   ", "  |"}
	output := ""
	if input == 2 {
		output += " _ \n"
		output += " _|\n"
		output += "|_"
	} else {
		output += top[1] + "\n"
		output += mid[1] + "\n"
		output += bot[1]
	}
	return output
}
