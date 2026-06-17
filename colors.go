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

func FgMagenta(s string) string {
	return "\033[35m" + s + "\033[0m"
}

func FgCyan(s string) string {
	return "\033[36m" + s + "\033[0m"
}
