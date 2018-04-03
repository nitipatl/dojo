package weather

func digit(input int) string {

	top := []string{"   ", "   ", " _ "}
	mid := []string{"   ", "  |", " _|"}
	bot := []string{"   ", "  |", "|_ "}
	output := ""
	output += top[input] + "\n"
	output += mid[input] + "\n"
	output += bot[input]
	return output
}
