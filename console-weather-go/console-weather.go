package weather

func digit(input int) string {

	top := []string{"   ", "   ", " _ "}
	mid := []string{"   ", "  |", " _|"}
	bot := []string{"   ", "  |", "|_ "}
	output := ""
	if input == 2 {
		output += top[2] + "\n"
		output += mid[2] + "\n"
		output += bot[2]
	} else {
		output += top[1] + "\n"
		output += mid[1] + "\n"
		output += bot[1]
	}
	return output
}
