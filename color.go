package main

type Color string

const (
	Reset         Color = "\033[0m"
	Black         Color = "\033[30m"
	Red           Color = "\033[31m"
	Green         Color = "\033[32m"
	Yellow        Color = "\033[33m"
	Blue          Color = "\033[34m"
	Magenta       Color = "\033[35m"
	Cyan          Color = "\033[36m"
	White         Color = "\033[37m"
	BrightBlack   Color = "\033[90m"
	BrightRed     Color = "\033[91m"
	BrightGreen   Color = "\033[92m"
	BrightYellow  Color = "\033[93m"
	BrightBlue    Color = "\033[94m"
	BrightMagenta Color = "\033[95m"
	BrightCyan    Color = "\033[96m"
	BrightWhite   Color = "\033[97m"
)

func (c Color) Format(s string) string {
	return string(c) + s + string(Reset)
}
