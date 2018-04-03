package weather

func digit(input int) string {

	top := []string{" _ ", "   ", " _ ", " _ ", "   ", " _ ", " _ ", " _ ", " _ ", " _ "}
	mid := []string{"| |", "  |", " _|", " _|", "|_|", "|_ ", "|_ ", "  |", "|_|", "|_|"}
	bot := []string{"|_|", "  |", "|_ ", " _|", "  |", " _|", "|_|", "  |", "|_|", " _|"}

	output := ""
	output += top[input] + "\n"
	output += mid[input] + "\n"
	output += bot[input]
	return output
}
