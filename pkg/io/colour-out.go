package io

import "github.com/fatih/color"

type PrintStyle int

const (
	Red PrintStyle = iota
	Green
	Blue
)

func GetStyle(code PrintStyle) *color.Color {
	switch code {
	case Red:
		return color.New(color.BgHiRed, color.Bold)
	case Green:
		return color.New(color.BgHiGreen, color.Bold)
	case Blue:
		return color.New(color.BgHiBlue, color.Bold)
	default:
		return color.New(color.BgBlack, color.Bold)
	}
}

func Panic(err error) {
	c := GetStyle(Red)
	c.Print("PANIC")
	println("")
}
func Fatal(err error) {
	c := GetStyle(Red)
	c.Print("FATAL")
	Fatal(err)
}

func Println(message ...any) {
	println(message)
}

func Print(message string) {
	print(message)
}

func PrintColor(message string, st PrintStyle) {
	c := GetStyle(st)
	c.Print(" " + message + " ")
	println("")
}
