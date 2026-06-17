package tineecolor

func Bold(s string) string {
	return "\033[1m" + s + "\033[0m"
}

func Italic(s string) string {
	return "\033[3m" + s + "\033[0m"
}

// Works not everywhere
func Underline(s string) string {
	return "\033[4m" + s + "\033[0m"
}

// Works not everywhere
func Blink(s string) string {
	return "\033[5m" + s + "\033[0m"
}
