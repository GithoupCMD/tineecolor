package tineecolor

func Italic(s string) string {
	return "\033[3m" + s + "\033[0m"
}

func Bold(s string) string {
	return "\033[1m" + s + "\033[0m"
}
