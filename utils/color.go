package utils

const (
	RED   = "\x1B[31m"
	GREEN = "\x1B[32m"
	RESET = "\x1B[0m"
)

func green(text string) string {
	return GREEN + text + RESET
}

func red(text string) string {
	return RED + text + RESET
}
