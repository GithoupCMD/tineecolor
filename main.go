package tineecolor

func FgRed(s string) string {
	return "\033[31m" + s + "\033[0m"
}

func FgGreen(s string) string {
	return "\033[32m" + s + "\033[0m"
}

func FgYellow(s string) string {
	return "\033[33m" + s + "\033[0m"
}

func FgBlue(s string) string {
	return "\033[34m" + s + "\033[0m"
}

func Italic(s string) string {
	return "\033[3m" + s + "\033[0m"
}

func Bold(s string) string {
	return "\033[1m" + s + "\033[0m"
}
