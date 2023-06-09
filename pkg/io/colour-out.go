package io

import (
	"fmt"
	"github.com/fatih/color"
)

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

func Panic(err ...interface{}) {
	c := GetStyle(Red)
	c.Print("PANIC")
	fmt.Println(err)
}
func Fatal(err ...interface{}) {
	c := GetStyle(Red)
	c.Print("FATAL")
	Fatal(err)
}

func Println(message ...interface{}) {

	fmt.Println(message)
}

func Print(message ...interface{}) {
	fmt.Println(message)
}

func PrintColor(message string, st PrintStyle) {
	c := GetStyle(st)
	c.Print(" " + message + " ")
	fmt.Println("")
}
